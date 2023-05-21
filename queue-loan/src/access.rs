use std::sync::Arc;

use axum::{
    extract::FromRequestParts,
    http::{request::Parts, StatusCode},
};
use axum_extra::extract::CookieJar;
use jsonwebtoken::{DecodingKey, Validation};
use serde::{Deserialize, Serialize};

use crate::state::ServerState;

#[derive(Debug, Serialize, Deserialize)]
pub struct MaybeAccess {
    pub user_id: Option<i32>,
}

#[axum::async_trait]
impl FromRequestParts<Arc<ServerState>> for MaybeAccess {
    type Rejection = StatusCode;

    async fn from_request_parts(
        parts: &mut Parts,
        state: &Arc<ServerState>,
    ) -> Result<Self, Self::Rejection> {
        from_request_parts(parts, state).await
    }
}

async fn from_request_parts<'p>(
    parts: &mut Parts,
    state: &Arc<ServerState>,
) -> Result<MaybeAccess, StatusCode> {
    let jar = CookieJar::from_request_parts(parts, state).await.unwrap();
    let Some(token) = jar.get("token") else {
        return Ok(MaybeAccess { user_id: None })
    };

    let claims = jsonwebtoken::decode(
        token.value(),
        &DecodingKey::from_secret(&state.jwt),
        &Validation::default(),
    )
    .map_err(|_| StatusCode::BAD_REQUEST)?
    .claims;

    Ok(claims)
}
