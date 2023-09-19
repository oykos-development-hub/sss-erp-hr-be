CREATE TABLE IF NOT EXISTS tenders_in_organization_units (
    id serial PRIMARY KEY,
    position_in_organization_unit_id int,
    type int,
    organization_unit_id int,
    date_of_start date not null,
    date_of_end date,
    description text,
    serial_number text not null,
    number_of_vacant_seats int,
    file_id int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (type) REFERENCES tender_types (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (organization_unit_id) REFERENCES organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (position_in_organization_unit_id) REFERENCES job_positions_in_organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
);
