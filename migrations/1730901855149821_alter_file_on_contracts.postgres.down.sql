ALTER TABLE employee_contracts ALTER COLUMN file_id TYPE INTEGER USING (file_id->>0)::INTEGER;
