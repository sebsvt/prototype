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

-- create profiles table
create table profiles (
    profile_id serial primary key,
    avatar text,
    firstname varchar(255),
    surname varchar(255),
    gender varchar(50),
    phone varchar(20),
    date_of_birth date,
    user_ref int unique not null,
    constraint fk_user
      foreign key(user_ref)
      references users(user_id)
);

create table studios (
    studio_id serial primary key,
    subdomain varchar(255) not null unique,
    picture varchar(255),
    name varchar(255) not null,
    description text,
    address varchar(255),
    city varchar(100),
    zipcode varchar(20),
    state varchar(100),
    country varchar(100),
    owner_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table memberships (
    id serial primary key,
    studio_id int not null,
    member_id int not null,
    role varchar(50) not null,
    foreign key (studio_id) references studios(studio_id),
    foreign key (member_id) references users(user_id)
);
