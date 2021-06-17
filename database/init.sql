create schema if not exists book_api;

create table if not exists book_api.category(
    id bigserial,
    name varchar(255) not null unique
);