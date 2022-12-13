CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT user_pk PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    "password" TEXT NOT NULL
);

CREATE UNIQUE INDEX user_email_idx ON users(email);