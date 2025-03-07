CREATE TABLE movies (
    created_at TIMESTAMP NOT NULL,
    description TEXT NOT NULL,
    duration SMALLINT NOT NULL,
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    watch_url TEXT NOT NULL
);
