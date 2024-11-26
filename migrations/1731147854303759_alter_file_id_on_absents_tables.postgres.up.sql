ALTER TABLE employee_absents
ADD COLUMN file_ids int[];

UPDATE employee_absents
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE employee_absents
DROP COLUMN file_id;