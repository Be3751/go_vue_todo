-- drop table tasks if exists;

create table users (
    id varchar(128) primary key,
    uuid varchar(64) not null unique, 
    enc_pwd varchar(255) not null
);

create table tasks (
    id int primary key,
    content text,
    user_id varchar(128),
    foreign key (user_id) references users(id) 
);