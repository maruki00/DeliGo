create database deligo;


create table usertypes (
    id int primary key auto_increment,
    label varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table users(
    id int primary key not null auto_increment,
    user_name varchar(100) not null unique,
    full_name varchar(255) not null,
    email varchar(255) not null unique,
    address text not null,
    password text not null,
    type int not null default 2,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table auth(
    id int primary key primary key auto_increment,
    email varchar(255) not null unique,
    token text not null unique,
    user_id int not null,
    user_type int not null,
    user_level int not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table orders(
    id int primary key not null auto_increment,
    order_fingerprint varchar(20) not null
    costumer_id int not null,
    from_long float not null,
    from_lat  float not null,
    to_long   float not null, 
    to_lat    float not null, 
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table products(
    id int primary key not null auto_increment,
    label varchar(255) not null,
    price varchar(10) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table product_images (
    id int primary key auto_increment,
    image_path varchar(100) not null,
    product_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table couriers_orders(
    id int primary key not null auto_increment,
    order_fingerprint varchar(20) not null,
    courier_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table feedbacks(
    id int primary key not null auto_increment,
    rate int not null default 0,
    comment text not null,
    service_type varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table analytics_order(
    id int primary key not null auto_increment,
    order_id int not null,
    user_id int not null,
    deligo_id int not null,
    rate int not null,
    comment text not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table orders_tracking(
    id int primary key not null auto_increment,
    order_id int not null,
    status varchar(100) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table orders_couriers(
    id int primary key not null auto_increment,
    user_id int not null,
    order_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table notifications(
    id int primary key not null auto_increment,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table jobs(
    id int primary key not null auto_increment,
    obj varchar(226) not null,
    data text not null,
    attempts int default 1,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
