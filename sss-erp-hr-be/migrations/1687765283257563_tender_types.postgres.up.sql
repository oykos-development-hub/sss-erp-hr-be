CREATE TABLE IF NOT EXISTS tender_types (
    id serial PRIMARY KEY,
    title text NOT NULL,
    is_judge boolean not null,
    is_judge_president boolean not null,
    abbreviation text NOT NULL,
    description text,
    color text,
    value text,
    icon text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
