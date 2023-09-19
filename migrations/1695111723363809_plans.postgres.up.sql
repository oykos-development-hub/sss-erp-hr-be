CREATE TABLE IF NOT EXISTS plans (
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    year TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
