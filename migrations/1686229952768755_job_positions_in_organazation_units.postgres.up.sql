CREATE TABLE IF NOT EXISTS job_positions_in_organization_units (
  id serial PRIMARY KEY,
  systematization_id INT REFERENCES systematizations(id) ON DELETE CASCADE,
  parent_organization_unit_id INT REFERENCES organization_units(id) ON DELETE CASCADE,
  job_position_id INT REFERENCES job_positions(id) ON DELETE CASCADE,
  available_slots INT DEFAULT 0,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);
