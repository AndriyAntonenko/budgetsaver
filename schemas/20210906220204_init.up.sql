CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    "id" uuid not null unique default uuid_generate_v4(),
    "name" varchar(255) not null,
    "email" varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "salt" varchar(255) not null,
    "last_login_at" timestamp with time zone null,
    "created_at" timestamp with time zone default current_timestamp,
    "deleted_at" timestamp with time zone null
);

CREATE TABLE finance_group
(
    "id" uuid not null unique default uuid_generate_v4(),
    "name" varchar(255) not null,
    "description" varchar(255) not null,
    "created_at" timestamp with time zone default current_timestamp,
    "deleted_at" timestamp with time zone null
);

CREATE TYPE finance_group_role AS ENUM ('owner', 'admin', 'member');

CREATE TABLE users_finance_group
(
    "user_id" uuid not null,
    "group_id" uuid not null,
    "role" finance_group_role not null,
    "created_at" timestamp with time zone default current_timestamp,
    "deleted_at" timestamp with time zone null,

    PRIMARY KEY ("user_id", "group_id"),
    
    CONSTRAINT finance_group_fk
        FOREIGN KEY("group_id") 
            REFERENCES finance_group("id")
                ON DELETE NO ACTION,

    CONSTRAINT user_fk
        FOREIGN KEY("user_id") 
            REFERENCES users("id")
                ON DELETE NO ACTION
);

CREATE TABLE budget
(
    "id" uuid not null unique default uuid_generate_v4(),
    "finance_group_id" uuid not null,
    "name" varchar(255) not null,
    "description" text null,
    "creator" uuid not null,
    "created_at" timestamp with time zone default current_timestamp,
    "deleted_at" timestamp with time zone null,

    CONSTRAINT budget_finance_group_fk
        FOREIGN KEY("finance_group_id") 
            REFERENCES finance_group("id")
                ON DELETE NO ACTION,

    CONSTRAINT budget_creator_fk
        FOREIGN KEY("creator") 
            REFERENCES users("id")
                ON DELETE NO ACTION
);
