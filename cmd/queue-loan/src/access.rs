use axum::{
    extract::FromRequestParts,
    http::{request::Parts, StatusCode},
};
use axum_extra::extract::CookieJar;
use jsonwebtoken::{DecodingKey, Validation};
use serde::{Deserialize, Serialize};
use std::future::Future;

#[derive(Debug, Serialize, Deserialize)]
pub struct Access {
    pub user_id: i32,
}

#[axum::async_trait]
impl<S: Send + Sync> FromRequestParts<S> for Access {
    type Rejection = StatusCode;

    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        from_request_parts(parts, state).await
    }
}

fn from_request_parts<'p, S: Send + Sync>(
    parts: &'p mut Parts,
    state: &'p S,
) -> impl Future<Output = Result<Access, StatusCode>> + Send + 'p {
    async move {
        let jar = CookieJar::from_request_parts(parts, state).await.unwrap();
        let token = jar.get("token").ok_or(StatusCode::UNAUTHORIZED)?.value();

        let key = include_bytes!("jwt.key");
        let claims = jsonwebtoken::decode(
            token,
            &DecodingKey::from_secret(key),
            &Validation::default(),
        )
        .map_err(|_| StatusCode::BAD_REQUEST)?
        .claims;

        Ok(claims)
    }
}
