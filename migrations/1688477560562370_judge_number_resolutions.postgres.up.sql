CREATE TABLE IF NOT EXISTS judge_number_resolutions (
    id serial PRIMARY KEY,
    active BOOLEAN NOT NULL,
    serial_number TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
