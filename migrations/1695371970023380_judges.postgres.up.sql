CREATE TABLE IF NOT EXISTS judges (
    id serial PRIMARY KEY,
    user_profile_id INTEGER NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    organization_unit_id INTEGER NOT NULL REFERENCES organization_units(id) ON DELETE CASCADE,
    norm_id INTEGER NOT NULL REFERENCES user_norm_fulfilments(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
