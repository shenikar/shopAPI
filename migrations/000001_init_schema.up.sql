-- Расширение для UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Таблица Address
CREATE TABLE
    address (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        country TEXT NOT NULL,
        city TEXT NOT NULL,
        street TEXT NOT NULL
    );

-- Таблица Client
CREATE TABLE
    client (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        client_name TEXT NOT NULL,
        client_surname TEXT NOT NULL,
        birthday DATE NOT NULL,
        gender VARCHAR(6) NOT NULL CHECK (gender IN ('male', 'female')),
        registration_date TIMESTAMP NOT NULL DEFAULT now (),
        address_id UUID NOT NULL REFERENCES address (id) ON DELETE CASCADE
    );

-- Таблица supplier
CREATE TABLE
    supplier (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        name TEXT NOT NULL,
        address_id UUID NOT NULL REFERENCES address (id) ON DELETE CASCADE,
        phone_number VARCHAR(20) NOT NULL
    );

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
        last_update_date DATE NOT NULL,
        supplier_id UUID NOT NULL REFERENCES supplier (id) ON DELETE CASCADE,
        image_id UUID REFERENCES images (id) ON DELETE SET NULL
    );