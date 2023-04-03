use axum::extract::Path;

pub async fn by_book_id(Path(_book_id): Path<i32>) {}
