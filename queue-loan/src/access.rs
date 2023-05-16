use axum::{
    extract::FromRequestParts,
    http::{request::Parts, StatusCode},
};
use axum_extra::extract::CookieJar;
use jsonwebtoken::{DecodingKey, Validation};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct MaybeAccess {
    pub user_id: Option<i32>,
}

#[axum::async_trait]
impl<S: Send + Sync> FromRequestParts<S> for MaybeAccess {
    type Rejection = StatusCode;

    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        from_request_parts(parts, state).await
    }
}

async fn from_request_parts<'p, S: Send + Sync>(
    parts: &'p mut Parts,
    state: &'p S,
) -> Result<MaybeAccess, StatusCode> {
    let jar = CookieJar::from_request_parts(parts, state).await.unwrap();
    let Some(token) = jar.get("token") else {
        return Ok(MaybeAccess { user_id: None })
    };

    let key = include_bytes!("jwt.key");
    let claims = jsonwebtoken::decode(
        token.value(),
        &DecodingKey::from_secret(key),
        &Validation::default(),
    )
    .map_err(|_| StatusCode::BAD_REQUEST)?
    .claims;

    Ok(claims)
}
