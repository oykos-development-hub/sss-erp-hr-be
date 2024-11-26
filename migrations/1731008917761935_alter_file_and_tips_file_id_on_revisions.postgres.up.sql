-- migrate file_id
ALTER TABLE revisions
ADD COLUMN file_ids int[];

UPDATE revisions
SET file_ids = CASE 
                 WHEN file_id IS NOT NULL THEN ARRAY[file_id] 
                 ELSE NULL 
               END;

ALTER TABLE revisions
DROP COLUMN file_id;

-- migrate tips_file_id
ALTER TABLE revisions
ADD COLUMN tips_file_ids int[];

UPDATE revisions
SET tips_file_ids = CASE 
                 WHEN tips_file_id IS NOT NULL THEN ARRAY[tips_file_id] 
                 ELSE NULL 
               END;

ALTER TABLE revisions
DROP COLUMN tips_file_id;


-- migrate from json int array to int[] for employee_contracts
ALTER TABLE employee_contracts
ADD COLUMN file_ids int[];

UPDATE employee_contracts
SET file_ids = ARRAY(
    SELECT jsonb_array_elements_text(file_id)::int
);

ALTER TABLE employee_contracts
DROP COLUMN file_id;

-- migrate from json int array to int[] for evaluations
ALTER TABLE evaluations
ADD COLUMN file_ids int[];

UPDATE evaluations
SET file_ids = ARRAY(
    SELECT jsonb_array_elements_text(file_id)::int
);

ALTER TABLE evaluations
DROP COLUMN file_id;


-- migrate from json int array to int[] for employee_experiences
ALTER TABLE employee_experiences
ADD COLUMN file_ids int[];

UPDATE employee_experiences
SET file_ids = ARRAY(
    SELECT jsonb_array_elements_text(file_id)::int
);

ALTER TABLE employee_experiences
DROP COLUMN file_id;
