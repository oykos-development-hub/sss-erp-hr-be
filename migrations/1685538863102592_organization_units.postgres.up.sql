drop table if exists organization_units;

CREATE TABLE organization_units (
    id serial PRIMARY KEY,
    parent_id INTEGER REFERENCES organization_units(id) ON DELETE CASCADE,
    title text NOT NULL,
    abbreviation text,
    pib text,
    order_id int,
    number_of_judges SMALLINT DEFAULT 0,
    color text NULL,
    icon text NULL,
    address text NULL,
    description text NULL,
    folder_id INTEGER NULL,
    bank_accounts TEXT[],
    code TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
