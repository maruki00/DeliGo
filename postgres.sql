


drop table if exists usertypes;
create table usertypes (
    id uuid primary key,
    label varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

DROP TABLE IF EXISTS users;
create table users(
    id uuid primary key ,
    email varchar(255) not null,
    password varchar(255) not null,
    role varchar(255) not null,
    -- user_name varchar(100) not null,
    -- full_name varchar(255) not null,
    -- email varchar(255) not null ,
    -- address varchar(255) not null,
    -- password varchar(255) not null,
    -- user_type int not null default 2,
    -- status int not null default 2,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
);

drop table if exists roles;
create table roles (
    id uuid primary key,
    label varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists permissions;
create table permissions (
    id uuid primary key,
    label varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists menus;
create table menus (
    id uuid primary key,
    label varchar(100) not null,
    route varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists auths;
create table auths(
    id uuid primary key ,
    email varchar(255) not null unique,
    token text not null unique,
    user_id int not null,
    user_type int not null,
    user_level int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);


drop table if exists orders;
create table orders(
    id uuid primary key ,
    order_fingerprint varchar(26) not null,
    costumer_id int not null,
    couurier_id int not null,
    status int not null,

    -- from_long float not null,
    -- from_lat  float not null,
    -- to_long   float not null, 
    -- to_lat    float not null, 
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);


drop table if exists products;
create table products(
    id uuid primary key ,
    label varchar(255) not null,
    price varchar(10) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);


drop table if exists product_images;
create table product_images (
   id uuid primary key ,
    image_path varchar(100) not null,
    product_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists couriers_orders;
create table couriers_orders(
    id uuid primary key ,
    order_fingerprint varchar(20) not null,
    courier_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists feedbacks;
create table feedbacks(
   id uuid primary key ,
    rate int not null default 0,
    comment text not null,
    service_type varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists analytics_order;
create table analytics_order(
    id uuid primary key ,
    order_id int not null,
    user_id int not null,
    delivery_id int not null,
    rate int not null,
    comment text not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);


drop table if exists orders_couriers;
create table orders_couriers(
    id uuid primary key ,
    user_id int not null,
    order_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists orders_tracking;
create table orders_tracking(
    id uuid primary key ,
    order_id int not null,
    langt float not null,
    longt float not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);

drop table if exists notifications; 
create table notifications(
    id uuid primary key ,
    body text not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delete_at  timestamp default null
);
