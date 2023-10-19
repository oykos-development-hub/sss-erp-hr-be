CREATE TABLE IF NOT EXISTS revision_revisors (
    id serial PRIMARY KEY,
    revision_id INTEGER REFERENCES revisions(id) ON DELETE CASCADE,
    revisor_id INTEGER REFERENCES user_profiles(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
