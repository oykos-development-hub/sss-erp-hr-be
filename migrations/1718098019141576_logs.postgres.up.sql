CREATE TABLE IF NOT EXISTS logs (
    id serial PRIMARY KEY,
    operation VARCHAR(10),
    entity TEXT,
    old_state JSONB,
    new_state JSONB,
    user_id INTEGER,
    item_id INTEGER,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION log_changes() RETURNS TRIGGER AS $$
DECLARE
    user_id INTEGER;
BEGIN
    BEGIN
        SELECT current_setting('myapp.user_id')::INTEGER INTO user_id;
    EXCEPTION
        WHEN others THEN
            user_id := NULL;  
    END;

    IF TG_OP = 'INSERT' THEN
        INSERT INTO logs (operation, new_state, user_id, item_id, entity)
        VALUES ('INSERT', row_to_json(NEW)::jsonb, user_id, NEW.id, TG_TABLE_NAME);
        RETURN NEW;
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO logs (operation, old_state, new_state, user_id, item_id, entity)
        VALUES ('UPDATE', row_to_json(OLD)::jsonb, row_to_json(NEW)::jsonb, user_id, NEW.id, TG_TABLE_NAME);
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO logs (operation, old_state, user_id, item_id, entity)
        VALUES ('DELETE', row_to_json(OLD)::jsonb, user_id, OLD.id, TG_TABLE_NAME);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER organization_units_insert
AFTER INSERT ON organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER organization_units_update
AFTER UPDATE ON organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER organization_units_delete
AFTER DELETE ON organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER job_positions_insert
AFTER INSERT ON job_positions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER job_positions_update
AFTER UPDATE ON job_positions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER job_positions_delete
AFTER DELETE ON job_positions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_profiles_insert
AFTER INSERT ON user_profiles
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_profiles_update
AFTER UPDATE ON user_profiles
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_profiles_delete
AFTER DELETE ON user_profiles
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER systematizations_insert
AFTER INSERT ON systematizations
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER systematizations_update
AFTER UPDATE ON systematizations
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER systematizations_delete
AFTER DELETE ON systematizations
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_contracts_insert
AFTER INSERT ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_contracts_update
AFTER UPDATE ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_contracts_delete
AFTER DELETE ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_resolutions_insert
AFTER INSERT ON employee_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_resolutions_update
AFTER UPDATE ON employee_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_resolutions_delete
AFTER DELETE ON employee_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_absents_insert
AFTER INSERT ON employee_absents
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_absents_update
AFTER UPDATE ON employee_absents
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER employee_absents_delete
AFTER DELETE ON employee_absents
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER tenders_in_organization_units_insert
AFTER INSERT ON tenders_in_organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER tenders_in_organization_units_update
AFTER UPDATE ON tenders_in_organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER tenders_in_organization_units_delete
AFTER DELETE ON tenders_in_organization_units
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_norm_fulfilments_insert
AFTER INSERT ON user_norm_fulfilments
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_norm_fulfilments_update
AFTER UPDATE ON user_norm_fulfilments
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER user_norm_fulfilments_delete
AFTER DELETE ON user_norm_fulfilments
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER salaries_insert
AFTER INSERT ON salaries
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER salaries_update
AFTER UPDATE ON salaries
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER salaries_delete
AFTER DELETE ON salaries
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER judge_number_resolutions_insert
AFTER INSERT ON judge_number_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER judge_number_resolutions_update
AFTER UPDATE ON judge_number_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER judge_number_resolutions_delete
AFTER DELETE ON judge_number_resolutions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER plans_insert
AFTER INSERT ON plans
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER plans_update
AFTER UPDATE ON plans
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER plans_delete
AFTER DELETE ON plans
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revisions_insert
AFTER INSERT ON revisions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revisions_update
AFTER UPDATE ON revisions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revisions_delete
AFTER DELETE ON revisions
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revision_tips_insert
AFTER INSERT ON revision_tips
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revision_tips_update
AFTER UPDATE ON revision_tips
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER revision_tips_delete
AFTER DELETE ON revision_tips
FOR EACH ROW
EXECUTE FUNCTION log_changes();