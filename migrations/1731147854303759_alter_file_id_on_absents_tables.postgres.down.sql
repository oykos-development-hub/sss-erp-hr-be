ALTER TABLE employee_absents
ADD COLUMN file_id int,

UPDATE employee_absents
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

ALTER TABLE employee_absents
DROP COLUMN file_ids,