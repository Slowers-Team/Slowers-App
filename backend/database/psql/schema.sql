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
    type VARCHAR(20) NOT NULL,
    phone_number VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL UNIQUE,
    post_address VARCHAR(50),
    postal_code INTEGER NOT NULL check(postal_code between 0 and 99999),
    city VARCHAR(40) NOT NULL,
    notes VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS Memberships (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_email VARCHAR(50) REFERENCES Users(email),
    business_id INTEGER REFERENCES Businesses(id),
    designation VARCHAR(20)
);
