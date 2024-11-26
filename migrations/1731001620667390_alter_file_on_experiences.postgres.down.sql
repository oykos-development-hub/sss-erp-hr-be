ALTER TABLE employee_experiences 
ALTER COLUMN file_id TYPE integer 
USING CASE 
        WHEN file_id = '[]'::jsonb THEN NULL 
        ELSE (file_id->>0)::integer 
      END;
