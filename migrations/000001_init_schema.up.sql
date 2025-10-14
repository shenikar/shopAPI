-- Расширение для UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enum для гендера
CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');

-- Таблица Address
CREATE TABLE
    address (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        country TEXT NOT NULL,
        city TEXT NOT NULL,
        street TEXT NOT NULL
    );

-- Таблица Client
CREATE TABLE
    client (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        client_name TEXT NOT NULL,
        client_surname TEXT NOT NULL,
        birthday DATE NOT NULL,
        gender gender_enum NOT NULL,
        registration_date TIMESTAMPTZ NOT NULL DEFAULT now (),
        address_id UUID NOT NULL REFERENCES address (id) ON DELETE CASCADE
    );

CREATE INDEX idx_client_name_surname ON client (client_name, client_surname);
CREATE INDEX idx_client_address_id ON client (address_id);

-- Таблица supplier
CREATE TABLE
    supplier (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        address_id UUID NOT NULL REFERENCES address (id) ON DELETE CASCADE,
        phone_number VARCHAR(20) NOT NULL
    );

CREATE INDEX idx_supplier_address_id ON supplier (address_id);

-- Таблица images
CREATE TABLE
    images (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        image BYTEA NOT NULL
    );

-- Таблица product
CREATE TABLE
    product (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        category TEXT NOT NULL,
        price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
        available_stock INT NOT NULL CHECK (available_stock >= 0),
        last_update_date TIMESTAMPTZ NOT NULL,
        supplier_id UUID NOT NULL REFERENCES supplier (id) ON DELETE CASCADE,
        image_id UUID REFERENCES images (id) ON DELETE SET NULL
    );

CREATE INDEX idx_product_supplier_id ON product (supplier_id);
CREATE INDEX idx_product_image_id ON product (image_id);