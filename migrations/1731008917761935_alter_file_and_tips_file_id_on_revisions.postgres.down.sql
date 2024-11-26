-- Down migration for `revisions` table
-- 1. Add back `file_id` and `tips_file_id` as int columns
ALTER TABLE revisions
ADD COLUMN file_id int,
ADD COLUMN tips_file_id int;

-- 2. Populate `file_id` and `tips_file_id` from `file_ids` and `tips_file_ids` respectively, using the first element
UPDATE revisions
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

UPDATE revisions
SET tips_file_id = tips_file_ids[1]
WHERE tips_file_ids IS NOT NULL;

-- 3. Drop `file_ids` and `tips_file_ids` columns
ALTER TABLE revisions
DROP COLUMN file_ids,
DROP COLUMN tips_file_ids;


-- Down migration for `employee_contracts` table
-- 1. Add back `file_id` as jsonb
ALTER TABLE employee_contracts
ADD COLUMN file_id jsonb;

-- 2. Populate `file_id` with `file_ids` values, converting them back to a JSON array format
UPDATE employee_contracts
SET file_id = to_jsonb(file_ids)
WHERE file_ids IS NOT NULL;

-- 3. Drop `file_ids` column
ALTER TABLE employee_contracts
DROP COLUMN file_ids;


-- Down migration for `evaluations` table
-- 1. Add back `file_id` as jsonb
ALTER TABLE evaluations
ADD COLUMN file_id jsonb;

-- 2. Populate `file_id` with `file_ids` values, converting them back to a JSON array format
UPDATE evaluations
SET file_id = to_jsonb(file_ids)
WHERE file_ids IS NOT NULL;

-- 3. Drop `file_ids` column
ALTER TABLE evaluations
DROP COLUMN file_ids;


-- Down migration for `employee_experiences` table
-- 1. Add back `file_id` as jsonb
ALTER TABLE employee_experiences
ADD COLUMN file_id jsonb;

-- 2. Populate `file_id` with `file_ids` values, converting them back to a JSON array format
UPDATE employee_experiences
SET file_id = to_jsonb(file_ids)
WHERE file_ids IS NOT NULL;

-- 3. Drop `file_ids` column
ALTER TABLE employee_experiences
DROP COLUMN file_ids;
