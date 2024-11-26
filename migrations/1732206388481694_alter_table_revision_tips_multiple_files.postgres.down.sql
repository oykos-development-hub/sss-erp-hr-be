ALTER TABLE revision_tips
ADD COLUMN file_id int,

UPDATE revision_tips
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

ALTER TABLE revision_tips
DROP COLUMN file_ids,