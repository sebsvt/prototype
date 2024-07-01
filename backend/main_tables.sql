create database prototype;
use prototype;

drop table if exists `nodes`;
create table nodes (
    node_id int not null auto_increment,
    reference varchar(255) not null,
    order_id int not null,
    duration int not null,
    is_active tinyint(1) not null default '1',
    deployment_date datetime not null,
    primary key (node_id),
    unique (order_id),
    unique (reference)
);

drop table if exists `orders`;
create table orders (
    order_id int auto_increment primary key,
    customer_id int not null,
    product_sku varchar(255) not null,
    product_cost decimal(10, 2) not null,
    duration int not null,
    payment_id int not null,
    created_at datetime not null
);

drop table if exists `payments`;
create table payments (
    payment_id int auto_increment primary key,
    sender varchar(255) not null,
    receiver varchar(255) not null,
    amount decimal(10, 2) not null,
    is_verified boolean not null,
    transaction_ref varchar(255) not null,
    transaction_time_stamp varchar(255) not null,
    created_at datetime not null
);

drop table if exists `products`;
create table products (
    product_id int auto_increment primary key,
    sku varchar(255) not null,
    name varchar(255) not null,
    description text,
    price decimal(10, 2) not null,
    is_available boolean not null
);
