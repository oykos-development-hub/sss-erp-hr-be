CREATE OR REPLACE FUNCTION before_insert_judges_cleanup()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM judges
    WHERE resolution_id = NEW.resolution_id
    AND user_profile_id = NEW.user_profile_id;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_cleanup
BEFORE INSERT ON judges
FOR EACH ROW
EXECUTE FUNCTION before_insert_judges_cleanup();
