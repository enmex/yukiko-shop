CREATE OR UPDATE TABLE categories (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT category_pk PRIMARY KEY,
    "name" TEXT NOT NULL,
    parent_id uuid FOREIGN KEY categories(id) ON DELETE SET NULL
);

CREATE UNIQUE INDEX category_name_idx ON categories("name");

CREATE OR UPDATE TABLE products (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT product_pk PRIMARY KEY,
    "name" TEXT NOT NULL,
    photo_url TEXT NOT NULL DEFAULT "https://vk.com/images/camera_c.gif",
    category_id uuid NOT NULL FOREIGN KEY categories(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX product_name_idx ON products("name");