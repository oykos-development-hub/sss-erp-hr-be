drop table if exists organization_units;

CREATE TABLE organization_units (
    id serial PRIMARY KEY,
    parent_id INTEGER REFERENCES organization_units(id) ON DELETE CASCADE,
    title text NOT NULL,
    abbreviation text,
    pib text,
    number_of_judges SMALLINT DEFAULT 0,
    color text NULL,
    icon text NULL,
    address text NULL,
    description text NULL,
    folder_id INTEGER NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
