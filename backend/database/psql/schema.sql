CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_login TIMESTAMPTZ,
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
    business_name VARCHAR(50) NOT NULL UNIQUE,
    business_type VARCHAR(20) NOT NULL,
    business_phone_number VARCHAR(30) NOT NULL UNIQUE,
    business_email VARCHAR(30) NOT NULL UNIQUE,
    business_address VARCHAR(50),
    postal_code INTEGER NOT NULL check(postal_code between 0 and 99999),
    city VARCHAR(40) NOT NULL,
    notes VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS Memberships (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_email INTEGER REFERENCES Users,
    business_id INTEGER REFERENCES Businesses,
    designation VARCHAR(20)
);
