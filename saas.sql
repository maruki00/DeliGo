CREATE TABLE tenants (
    id uuid primary key,
    name varchat(100) not null,
    logo varchar(255) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
);
CREATE TABLE users (
    id uuid primary key,
    username varchar(32) not null,
    email varchar(255) not null,
    tenant_id varchart(32) not null,
    password varchar(255) not null,
    password_changed_at timestamp default null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
);
CREATE TABLE policies (
    id uuid primary key,
    name varchar(255) not null,
    group_id varchar(32) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
);
CREATE TABLE permissions (
    id uuid primary key,
    name varchar(255) not null,
    description text not null,
    policy_id varchar(255) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
)
CREATE TABLE profile (
    id uuid primary key,
    user_id varchar(32) not null,
    avatar varchar(255) not null,
    group_id varrchar(32) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
)
CREATE TABLE groups (
    id uuid primary key,
    name varchar(255) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
)
CREATE TABLE group_policies (
    id uuid primary key,
    group_id varchar(32) not null,
    policy_id varchar(32) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
)
CREATE TABLE user_groups (
    id uuid primary key,
    user_id  varchar(32) not null,
    group_id varchar(32) not null,
    deleted_at timestamp default now(),
    updated_at timestamp not null default now(),
    created_at timestamp not null default now()
)
