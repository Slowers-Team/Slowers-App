CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_login TIMESTAMPTZ,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE,
    is_active BOOLEAN,
    is_admin BOOLEAN
);

CREATE TABLE IF NOT EXISTS Businesses (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(50) UNIQUE,
    type VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS Memberships (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_email INTEGER REFERENCES Users,
    business_id INTEGER REFERENCES Businesses,
    designation VARCHAR(20)
);
