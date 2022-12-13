CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR UPDATE TABLE roles (
    "name" VARCHAR NOT NULL CONSTRAINT role_pk PRIMARY KEY
);

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT user_pk PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "role" VARCHAR(15) NOT NULL FOREIGN KEY roles("name")
);

CREATE UNIQUE INDEX user_email_idx ON users(email);