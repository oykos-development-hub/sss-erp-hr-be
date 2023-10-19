CREATE TABLE IF NOT EXISTS revisions_in_organization_units (
    id serial PRIMARY KEY,
    revision_id INTEGER REFERENCES revisions(id) ON DELETE CASCADE,
    organization_unit_id INTEGER REFERENCES organization_units(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
