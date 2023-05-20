use axum::{http::StatusCode, response::IntoResponse};
use thiserror::Error;
use tracing::error;

#[derive(Debug, Error)]
pub enum Error {
    #[error("Sqlx error: {0}")]
    SqlxError(#[from] sqlx::Error),
    #[error("JWT error: {0}")]
    JwtError(#[from] jsonwebtoken::errors::Error),
}

pub type Result<T> = std::result::Result<T, Error>;

impl IntoResponse for Error {
    fn into_response(self) -> axum::response::Response {
        let status = match self {
            Error::SqlxError(_) => StatusCode::INTERNAL_SERVER_ERROR,
            Error::JwtError(_) => StatusCode::BAD_REQUEST,
        };

        if status.is_server_error() {
            error!("Error: {self}");
        }

        status.into_response()
    }
}
