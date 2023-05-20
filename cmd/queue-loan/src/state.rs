use axum::extract::State;
use sqlx::PgPool;
use std::sync::Arc;

pub struct ServerState {
    pub postgres: PgPool,
}

pub type StateExtract = State<Arc<ServerState>>;
