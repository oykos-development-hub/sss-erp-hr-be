ALTER TABLE employee_educations
ADD COLUMN file_id int,

UPDATE employee_educations
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

ALTER TABLE employee_educations
DROP COLUMN file_ids,