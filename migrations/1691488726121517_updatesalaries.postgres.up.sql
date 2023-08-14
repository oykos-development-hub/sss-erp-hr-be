-- Add the organization_unit_id column
ALTER TABLE salaries
ADD COLUMN organization_unit_id INTEGER;

-- Add the foreign key constraint
ALTER TABLE salaries
ADD CONSTRAINT fk_organization_unit
FOREIGN KEY (organization_unit_id) REFERENCES organization_units (id)
ON UPDATE CASCADE ON DELETE CASCADE;
