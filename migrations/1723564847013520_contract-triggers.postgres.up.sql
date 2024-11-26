CREATE OR REPLACE FUNCTION set_active_false()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE employee_contracts
    SET active = FALSE
    WHERE user_profile_id = NEW.user_profile_id
      AND active = TRUE;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_active_false_trigger
BEFORE INSERT OR UPDATE ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION set_active_false();
