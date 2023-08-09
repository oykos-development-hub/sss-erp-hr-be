CREATE TABLE IF NOT EXISTS job_positions_in_organization_units (
  id serial PRIMARY KEY,
  systematization_id INT REFERENCES systematizations(id) ON DELETE CASCADE,
  parent_organization_unit_id INT REFERENCES organization_units(id) ON DELETE CASCADE,
  job_position_id INT REFERENCES job_positions(id) ON DELETE CASCADE,
  parent_job_position_id INT NULL REFERENCES job_positions_in_organization_units(id) ON DELETE CASCADE,
  system_permission_id INT NULL,
  description text,
  serial_number text NOT NULL,
  available_slots INT DEFAULT 0,
  requirements text NULL,
  icon text,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);
