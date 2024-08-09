-- drop tables if they exist
drop table if exists profiles;
drop table if exists users;
drop table if exists studios;
drop table if exists memberships;
-- create users table
create table users (
    user_id serial primary key,
    email varchar(255) unique not null,
    hashed_password varchar(255) not null,
    is_active boolean default true,
    is_verified boolean default false,
    last_signed_in timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);
