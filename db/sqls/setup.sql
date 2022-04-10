-- drop table tasks if exists;

create table tasks (
    id serial primary key,
    content text
);

create table users (
    id serial primary key,
    uuid varchar(64) not null unique,
    enc_pwd varchar(255) not null
)

create table sessions (
    id serial primary key,
    uuid varchar(64) not null unique
)