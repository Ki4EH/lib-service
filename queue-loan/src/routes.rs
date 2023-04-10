use crate::{
    access::Access,
    error::Result,
    models::{Record, Status},
    state::ServerState,
};
use axum::extract::{Path, State};
use chrono::{Duration, Utc};
use std::sync::Arc;

// GET /status/:book_id
pub async fn status_by_book_id(
    State(state): State<Arc<ServerState>>,
    Path(book_id): Path<i32>,
) -> Result<&'static str> {
    let mut tx = state.postgres.begin().await?;
    let records = sqlx::query_as!(
        Record,
        r#"SELECT user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1"#,
        book_id
    )
    .fetch_all(&mut tx)
    .await?;

    let status = if records.is_empty() {
        "available"
    } else {
        "unavailable"
    };

    Ok(status)
}

// POST /loan/:book_id
pub async fn loan_by_book_id(
    State(state): State<Arc<ServerState>>,
    Path(book_id): Path<i32>,
    access: Access,
) -> Result<()> {
    let mut tx = state.postgres.begin().await?;
    let records = sqlx::query_as!(
        Record,
        r#"SELECT user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1 ORDER BY id ASC"#,
        book_id
    ).fetch_all(&mut tx).await?;

    if records.iter().any(|r| r.user_id == access.user_id) {
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
            access.user_id,
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
            access.user_id,
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
    State(state): State<Arc<ServerState>>,
    Path(book_id): Path<i32>,
    access: Access,
) -> Result<()> {
    let mut tx = state.postgres.begin().await?;

    let queue = sqlx::query_as!(
        Record,
        r#"SELECT user_id, book_id, date, status as "status: Status" FROM queue WHERE book_id = $1 ORDER BY id ASC"#,
        book_id
    )
    .fetch_all(&mut tx).await?;

    if queue.is_empty() {
        return Ok(());
    }

    sqlx::query!(
        "DELETE FROM queue WHERE user_id = $1 AND book_id = $2",
        access.user_id,
        book_id
    )
    .execute(&mut tx)
    .await?;

    // If removed user is not first do nothing.
    let first = queue.first().unwrap();
    if first.user_id != access.user_id {
        return Ok(());
    }

    // Otherwise set next user to PendingBorrow status.
    if let Some(next) = queue.get(1) {
        sqlx::query!(
            "UPDATE queue SET status = $1 WHERE id = $2",
            Status::PendingBorrow as i32,
            next.user_id
        )
        .execute(&mut tx)
        .await?;
    }

    tx.commit().await?;
    Ok(())
}
