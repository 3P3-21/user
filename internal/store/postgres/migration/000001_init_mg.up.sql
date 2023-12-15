-- users db

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email TEXT UNIQUE,
    password_hash TEXT,
    first_name TEXT,
    last_name TEXT,
    role_id integer REFERENCES roles (id),
    created_at timestamp DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS roles (
    id serial PRIMARY KEY,
    name TEXT,
);

CREATE TABLE IF NOT EXISTS sessions (
    id serial PRIMARY KEY,
    session TEXT UNIQUE,
    expires_at timestamp,
    created_at timestamp DEFAULT NOW()
);
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';

GRANT SELECT, INSERT, UPDATE, DELETE ON database_name.* TO 'username'@'localhost';
