use crate::{
    access::MaybeAccess,
    error::{Error, Result},
    models::{Order, Record, Status},
    state::StateExtract,
};
use axum::{
    extract::{Path, State},
    Json,
};
use chrono::{Duration, Utc};
use std::mem::transmute;

// GET /orders
pub async fn orders(State(state): StateExtract, access: MaybeAccess) -> Result<Json<Vec<Order>>> {
    let user_id = access.user_id.ok_or(Error::Unauthorized)?;
    let mut tx = state.postgres.begin().await?;
    let records = sqlx::query!(
        "SELECT queue.book_id, book.name as book_name, author.name as book_author, queue.date, queue.status
        FROM queue
        INNER JOIN book ON book.id = queue.book_id
        INNER JOIN author ON author.id = book.author_id
        WHERE queue.user_id = $1",
        user_id
    )
    .fetch_all(&mut tx)
    .await?;

    Ok(records
        .into_iter()
        .enumerate()
        .map(|(i, r)| Order {
            book_id: r.book_id,
            // SAFETY: Status is repr(i32), and only inserted based on the Status enum.
            status: unsafe { transmute(r.status) },
            position: i as u32,
            date: r.date,
            book_name: r.book_name,
            book_author: r.book_author,
        })
        .collect::<Vec<_>>()
        .into())
}

// GET /status/:book_id
pub async fn status_by_book_id(
    State(state): StateExtract,
    Path(book_id): Path<i32>,
    access: MaybeAccess,
) -> Result<String> {
    let mut tx = state.postgres.begin().await?;
    let records = sqlx::query_as!(
        Record,
        r#"SELECT id, user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1"#,
        book_id
    )
    .fetch_all(&mut tx)
    .await?;

    let status = if records.is_empty() {
        "available".to_owned()
    } else if let Some(i) = records
        .iter()
        .position(|r| Some(r.user_id) == access.user_id)
    {
        i.to_string()
    } else {
        "unavailable".to_owned()
    };

    Ok(status)
}

// POST /loan/:book_id
pub async fn loan_by_book_id(
    State(state): StateExtract,
    Path(book_id): Path<i32>,
    access: MaybeAccess,
) -> Result<()> {
    let user_id = access.user_id.ok_or(Error::Unauthorized)?;

    let mut tx = state.postgres.begin().await?;
    let records = sqlx::query_as!(
        Record,
        r#"SELECT id, user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1 ORDER BY id ASC"#,
        book_id
    ).fetch_all(&mut tx).await?;

    if records.iter().any(|r| r.user_id == user_id) {
        // Skip if user already queued.
        return Ok(());
    }

    let is_first = records.is_empty();
    if is_first {
        let deadline = Utc::now() + Duration::days(7);
        sqlx::query!(
            "
            INSERT INTO queue(user_id, book_id, date, status)
            VALUES ($1, $2, $3, $4)",
            user_id,
            book_id,
            Some(deadline),
            Status::PendingBorrow as i32,
        )
        .execute(&mut tx)
        .await?;
    } else {
        sqlx::query!(
            "
            INSERT INTO queue(user_id, book_id, status)
            VALUES ($1, $2, $3)",
            user_id,
            book_id,
            Status::Queued as i32,
        )
        .execute(&mut tx)
        .await?;
    }

    tx.commit().await?;

    Ok(())
}

// POST /cancel/:book_id
pub async fn cancel_by_book_id(
    State(state): StateExtract,
    Path(book_id): Path<i32>,
    access: MaybeAccess,
) -> Result<()> {
    let user_id = access.user_id.ok_or(Error::Unauthorized)?;
    let mut tx = state.postgres.begin().await?;

    let queue = sqlx::query_as!(
        Record,
        r#"SELECT id, user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1 ORDER BY id ASC"#,
        book_id
    )
    .fetch_all(&mut tx).await?;

    if queue.is_empty() {
        return Ok(());
    }

    sqlx::query!(
        "DELETE FROM queue WHERE user_id = $1 AND book_id = $2",
        user_id,
        book_id
    )
    .execute(&mut tx)
    .await?;

    // If removed user is not first do nothing.
    let first = queue.first().unwrap();
    if first.user_id != user_id {
        tx.commit().await?;
        return Ok(());
    }

    // Otherwise set next user to PendingBorrow status.
    if let Some(next) = queue.get(1) {
        sqlx::query!(
            "UPDATE queue SET status = $1 WHERE id = $2",
            Status::PendingBorrow as i32,
            next.id
        )
        .execute(&mut tx)
        .await?;
    }

    tx.commit().await?;
    Ok(())
}