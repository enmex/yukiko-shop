CREATE TABLE cart_products (
    id uuid NOT NULL DEFAULT uuid_generate_v4() CONSTRAINT cart_pk PRIMARY KEY,
    product_id uuid NOT NULL REFERENCES products(id) ON DELETE SET NULL,
    customer_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "name" TEXT NOT NULL,
    photo_url TEXT NOT NULL DEFAULT 'https://vk.com/images/camera_c.gif',
    price DOUBLE PRECISION NOT NULL,
    quantity INT NOT NULL
);

CREATE UNIQUE INDEX cart_product_name_idx ON cart_products("name");