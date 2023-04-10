use axum::{
    extract::FromRequestParts,
    http::{header::AUTHORIZATION, request::Parts, StatusCode},
};
use jsonwebtoken::{DecodingKey, Validation};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Access {
    pub user_id: i32,
}

#[axum::async_trait]
impl<S> FromRequestParts<S> for Access {
    type Rejection = StatusCode;

    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        let (_, value) = parts
            .headers
            .iter()
            .find(|(name, value)| **name == AUTHORIZATION)
            .ok_or(StatusCode::UNAUTHORIZED)?;

        let key = include_bytes!("jwt.key");
        let claims = jsonwebtoken::decode(
            value.to_str().map_err(|_| StatusCode::BAD_REQUEST)?,
            &DecodingKey::from_secret(key),
            &Validation::default(),
        )
        .map_err(|_| StatusCode::BAD_REQUEST)?
        .claims;

        Ok(claims)
    }
}
