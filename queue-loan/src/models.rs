use serde::{Deserialize, Serialize};
use sqlx::{
    types::chrono::{DateTime, Utc},
    FromRow, Type,
};

#[derive(Debug, FromRow)]
pub struct Record {
    pub id: i32,
    pub user_id: i32,
    pub book_id: i32,
    pub status: Status,
    pub date: Option<DateTime<Utc>>,
}

#[derive(Debug, Serialize, Deserialize, Type)]
#[repr(i32)]
pub enum Status {
    PendingBorrow,
    Acquried,
    Queued,
}

#[derive(Debug, FromRow)]
pub struct User {
    pub id: i32,
    pub login: String,
    pub email: String,
    pub password_hash: String,
    pub flags: String,
    pub confirm_token: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Order {
    pub book_id: i32,
    pub status: Status,
    pub position: u32,
    pub date: Option<DateTime<Utc>>,
}
