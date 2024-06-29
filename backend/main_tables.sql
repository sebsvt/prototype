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
    order_id int not null,
    customer_id int not null,
    product_sku varchar(255) not null,
    product_cost decimal(10, 2) not null,
    duration int not null,
    created_at datetime not null,
    primary key (order_id)
);
