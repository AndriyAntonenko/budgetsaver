CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    "id" uuid not null unique default uuid_generate_v4(),
    "name" varchar(255) not null,
    "email" varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "salt" varchar(255) not null
);

CREATE TABLE "group" 
(
    "id" uuid not null unique default uuid_generate_v4(),
    "name" varchar(255) not null,
    "type" varchar(255) not null
);

CREATE TABLE users_groups
(
    "user_id" uuid references users (id) not null,
    "group_id" uuid references "group" (id) not null 
);

CREATE TABLE budget
(
    "id" uuid not null unique default uuid_generate_v4(),
    "group_id" uuid references "group" (id) not null,
    "name" varchar(255) not null unique,
    "description" text
);

ALTER TABLE users_groups ADD PRIMARY KEY ("user_id", "group_id");
