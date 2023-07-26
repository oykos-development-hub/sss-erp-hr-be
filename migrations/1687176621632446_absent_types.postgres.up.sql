CREATE TABLE IF NOT EXISTS absent_types (
    id serial PRIMARY KEY,
    parent_id INT REFERENCES absent_types (id),
    title VARCHAR (255) NOT NULL,
    abbreviation VARCHAR (8) NOT NULL,
    relocation BOOLEAN NOT NULL,
    accounting_days_off BOOLEAN NOT NULL,
    description text,
    color VARCHAR (32),
    icon text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);
