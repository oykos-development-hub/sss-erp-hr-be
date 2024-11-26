ALTER TABLE employee_resolutions
ADD COLUMN file_ids int[];

UPDATE employee_resolutions
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL AND file_id != 0 THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE employee_resolutions
DROP COLUMN file_id;
