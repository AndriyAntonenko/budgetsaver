CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    "id" uuid not null unique default uuid_generate_v4(),
    "name" varchar(255) not null,
    "email" varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "salt" varchar(255) not null
);
