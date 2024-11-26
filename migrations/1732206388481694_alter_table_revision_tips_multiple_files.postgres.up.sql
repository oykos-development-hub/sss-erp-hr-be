ALTER TABLE revision_tips
ADD COLUMN file_ids int[];

UPDATE revision_tips
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE revision_tips
DROP COLUMN file_id;