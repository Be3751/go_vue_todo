-- drop table tasks if exists;

create table tasks (
    id serial primary key,
    content text
);

create table users (
    id serial primary key,
    enc_pwd varchar(256) not null
)