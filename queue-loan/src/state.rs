use sqlx::PgPool;

pub struct ServerState {
    pub postgres: PgPool,
}
