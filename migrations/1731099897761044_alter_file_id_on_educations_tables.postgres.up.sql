ALTER TABLE employee_educations
ADD COLUMN file_ids int[];

UPDATE employee_educations
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL AND file_id != 0 THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE employee_educations
DROP COLUMN file_id;