DROP TRIGGER IF EXISTS trg_delete_inactive_employees ON employees_in_organization_units;

CREATE OR REPLACE FUNCTION delete_inactive_employees()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the related systematization is active = 2
    IF EXISTS (
        SELECT 1
        FROM job_positions_in_organization_units jpou
        JOIN systematizations s ON jpou.systematization_id = s.id
        WHERE jpou.id = NEW.position_in_organization_unit_id
        AND s.active = 2
    ) THEN
        -- Delete employees with the same user_profile_id and systematization active = 2
        DELETE FROM employees_in_organization_units
        WHERE user_profile_id = NEW.user_profile_id
        AND position_in_organization_unit_id IN (
            SELECT jpou.id
            FROM job_positions_in_organization_units jpou
            JOIN systematizations s ON jpou.systematization_id = s.id
            WHERE s.active = 2 
        );
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger on the employees_in_organization_units table
CREATE TRIGGER trigger_delete_inactive_employees
AFTER INSERT ON employees_in_organization_units
FOR EACH ROW
EXECUTE FUNCTION delete_inactive_employees();
