CREATE TABLE IF NOT EXISTS judge_number_resolution_organization_units (
    id serial PRIMARY KEY,
    resolution_id INTEGER NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    number_of_judges INTEGER NOT NULL,
    number_of_presidents INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    FOREIGN KEY (resolution_id) REFERENCES judge_number_resolutions (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (organization_unit_id) REFERENCES organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
)