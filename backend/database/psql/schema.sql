CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_login TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    is_active BOOLEAN,
    is_admin BOOLEAN
);

CREATE TABLE IF NOT EXISTS Businesses (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(50) NOT NULL UNIQUE,
    business_id_code VARCHAR(50),
    type VARCHAR(20) NOT NULL,
    phone_number VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL UNIQUE,
    address VARCHAR(50),
    postal_code VARCHAR(5) NOT NULL,
    city VARCHAR(40) NOT NULL,
    additional_info VARCHAR(1500) NOT NULL
);

CREATE TABLE IF NOT EXISTS Memberships (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_email VARCHAR(50),
    business_id INTEGER REFERENCES Businesses(id),
    designation VARCHAR(20)
);
