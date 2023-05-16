use axum::{
    routing::{get, post},
    Router, Server,
};
use eyre::eyre;
use sqlx::PgPool;
use state::ServerState;
use std::{env::var, sync::Arc};
use tower_http::cors::{AllowOrigin, CorsLayer};
use tracing::{debug, log::info};

mod access;
mod error;
mod models;
mod routes;
mod state;

#[tokio::main]
async fn main() -> eyre::Result<()> {
    run().await
}

async fn create_state() -> eyre::Result<Arc<ServerState>> {
    let database_url =
        var("DATABASE_URL").map_err(|_| eyre!("Missing DATABASE_URL environment variable"))?;
    debug!("Database url: {database_url}");

    let pool = PgPool::connect(&database_url).await?;
    info!("Successfully connected to the database.");

    Ok(ServerState { postgres: pool }.into())
}

async fn run() -> eyre::Result<()> {
    dotenvy::dotenv()?;
    tracing_subscriber::fmt::init();

    #[cfg(debug_assertions)]
    let cors = CorsLayer::permissive();
    #[cfg(not(debug_assertions))]
    let cors = CorsLayer::new();

    let state = create_state().await?;
    let service_root = Router::new()
        .route("/api/queue/status/:book_id", get(routes::status_by_book_id))
        .route("/api/queue/loan/:book_id", post(routes::loan_by_book_id))
        .route(
            "/api/queue/cancel/:book_id",
            post(routes::cancel_by_book_id),
        )
        .layer(cors)
        .with_state(state);

    Server::bind(&"0.0.0.0:7000".parse()?)
        .serve(service_root.into_make_service())
        .await?;

    Ok(())
}
