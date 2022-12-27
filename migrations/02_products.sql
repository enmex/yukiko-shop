CREATE TABLE categories (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT category_pk PRIMARY KEY,
    "name" TEXT NOT NULL,
    photo_url TEXT NOT NULL DEFAULT 'https://vk.com/images/camera_c.gif',
    parent_category uuid REFERENCES categories(id) ON DELETE SET NULL
);

CREATE UNIQUE INDEX category_name_idx ON categories("name");

CREATE TABLE products (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT product_pk PRIMARY KEY,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    photo_url TEXT NOT NULL DEFAULT 'https://vk.com/images/camera_c.gif',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    category_id uuid NOT NULL REFERENCES categories(id) ON DELETE SET NULL
);

CREATE UNIQUE INDEX product_name_idx ON products("name");