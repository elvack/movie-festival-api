CREATE TABLE admins (
    created_at TIMESTAMP NOT NULL,
    email TEXT UNIQUE NOT NULL,
    encrypted_password TEXT NOT NULL,
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
