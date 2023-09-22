CREATE TABLE IF NOT EXISTS judges (
    id serial PRIMARY KEY,
    user_profile_id INTEGER NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    organization_unit_id INTEGER NOT NULL REFERENCES organization_units(id) ON DELETE CASCADE,
    resolution_id INTEGER NOT NULL REFERENCES judge_number_resolutions(id) ON DELETE CASCADE,
    is_president BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
