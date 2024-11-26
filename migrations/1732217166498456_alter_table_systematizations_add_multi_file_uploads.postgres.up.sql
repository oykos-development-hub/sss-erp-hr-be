ALTER TABLE systematizations
ADD COLUMN file_ids int[];

UPDATE systematizations
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL AND file_id != 0 THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE systematizations
DROP COLUMN file_id;
