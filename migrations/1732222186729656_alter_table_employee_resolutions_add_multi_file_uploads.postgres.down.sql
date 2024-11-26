ALTER TABLE employee_resolutions
ADD COLUMN file_id int,

UPDATE employee_resolutions
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

ALTER TABLE employee_resolutions
DROP COLUMN file_ids,