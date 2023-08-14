-- Drop the foreign key constraint
ALTER TABLE salaries
DROP CONSTRAINT fk_organization_unit;

-- Drop the organization_unit_id column
ALTER TABLE salaries
DROP COLUMN organization_unit_id;
