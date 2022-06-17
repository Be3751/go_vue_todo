-- drop table tasks if exists;

create table users (
    id varchar(128) primary key,
    uuid varchar(64) unique, 
    enc_pwd varchar(255) not null
);

create table tasks (
    id serial not null primary key,
    content text,
    deadline date,
    created_at date,
    updated_at date,
    user_id varchar(128),
    foreign key (user_id) references users(id) 
);