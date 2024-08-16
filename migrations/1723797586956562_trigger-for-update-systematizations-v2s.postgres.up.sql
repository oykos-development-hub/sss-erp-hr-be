CREATE OR REPLACE FUNCTION delete_inactive_employees()
RETURNS TRIGGER AS $$
BEGIN

    IF EXISTS (
        SELECT 1
        FROM job_positions_in_organization_units jpou
        JOIN systematizations s ON jpou.systematization_id = s.id
        WHERE jpou.id = NEW.position_in_organization_unit_id
        AND s.active = 2
    ) THEN
        DELETE FROM employees_in_organization_units
        WHERE user_profile_id = NEW.user_profile_id
        AND position_in_organization_unit_id IN (
            SELECT jpou.id
            FROM job_positions_in_organization_units jpou
            JOIN systematizations s ON jpou.systematization_id = s.id
            WHERE s.active = 2 
        );

        DELETE FROM judges
        WHERE user_profile_id = NEW.user_profile_id
        AND EXISTS (
            SELECT 1
            FROM judge_number_resolutions r
            WHERE r.id = judges.resolution_id and r.active = true
        );
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
