ALTER TABLE employee_experiences 
ALTER COLUMN file_id TYPE jsonb 
USING CASE 
        WHEN file_id IS NULL THEN '[]'::jsonb 
        ELSE jsonb_build_array(file_id) 
      END;
