use axum::{routing::get, Router, Server};
use eyre::{eyre, Result};
use sqlx::PgPool;
use state::ServerState;
use std::{env::var, sync::Arc};
use tracing::{debug, log::info};

mod models;
mod state;
mod status;

#[tokio::main]
async fn main() -> Result<()> {
    run().await
}

async fn create_state() -> Result<Arc<ServerState>> {
    let database_url =
        var("DATABASE_URL").map_err(|_| eyre!("Missing DATABASE_URL environment variable"))?;
    debug!("Database url: {database_url}");

    let pool = PgPool::connect(&database_url).await?;
    info!("Successfully connected to the database.");

    Ok(ServerState { postgres: pool }.into())
}

async fn run() -> Result<()> {
    dotenvy::dotenv()?;
    tracing_subscriber::fmt::init();

    let _state = create_state().await?;
    let service_root = Router::new().route("/api/queue/status", get(status::by_book_id));

    Server::bind(&"0.0.0.0:7000".parse()?)
        .serve(service_root.into_make_service())
        .await?;

    Ok(())
}
