ALTER TABLE systematizations
ADD COLUMN file_id int,

UPDATE systematizations
SET file_id = file_ids[1]
WHERE file_ids IS NOT NULL;

ALTER TABLE systematizations
DROP COLUMN file_ids,