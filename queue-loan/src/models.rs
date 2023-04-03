#[derive(Debug)]
pub struct Book {
    pub id: i32,
    pub name: String,
    pub author_id: i32,
    pub genre_id: i32,
    pub isbn: String,
}

#[derive(Debug)]
pub struct Author {
    pub id: i32,
    pub name: String,
}

#[derive(Debug)]
pub struct Catalog {
    pub id: i32,
    pub book_id: i32,
    pub count: i32,
}

#[derive(Debug)]
pub struct User {
    pub id: i32,
    pub login: String,
    pub email: String,
    pub password_hash: String,
    pub flags: String,
    pub confirm_token: Option<String>,
}
