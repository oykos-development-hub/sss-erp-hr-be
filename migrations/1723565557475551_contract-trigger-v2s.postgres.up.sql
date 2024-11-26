DROP TRIGGER IF EXISTS set_active_false_trigger ON employee_contracts;

DROP FUNCTION IF EXISTS set_active_false();

CREATE OR REPLACE FUNCTION activate_new_contract()
RETURNS TRIGGER AS $$
DECLARE
    old_active BOOLEAN;
BEGIN
    PERFORM 1 FROM employee_contracts
    WHERE user_profile_id = NEW.user_profile_id AND id <> NEW.id AND active = TRUE;

    IF FOUND THEN
        UPDATE employee_contracts
        SET active = FALSE
        WHERE user_profile_id = NEW.user_profile_id AND id <> NEW.id;
    END IF;

    NEW.active = TRUE;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER activate_new_contract_trigger
BEFORE INSERT OR UPDATE ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION activate_new_contract();
