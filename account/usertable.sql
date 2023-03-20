create table users(
    id serial primary key,
    login varchar(128) unique not null,
    email varchar(128) unique not null,
    password_hash varchar(128) not null,
    flags varchar(128),
    confirm_token varchar(256) null,
);
\! chcp 1251