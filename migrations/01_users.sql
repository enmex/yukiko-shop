CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE access_type AS ENUM (
    'ADMIN',
    'MODERATOR',
    'CUSTOMER'
);

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT user_pk PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    "password" TEXT NOT NULL,
    access_type access_type NOT NULL DEFAULT 'CUSTOMER'
);

CREATE UNIQUE INDEX user_email_idx ON users(email);