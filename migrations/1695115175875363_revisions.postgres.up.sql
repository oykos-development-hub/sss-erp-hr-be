CREATE TABLE IF NOT EXISTS revisions (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    plan_id INTEGER NOT NULL REFERENCES plans(id) ON DELETE CASCADE, 
    serial_number TEXT NOT NULL,
    date_of_revision DATE NOT NULL,
    revision_quartal TEXT NOT NULL,
    internal_revision_subject INTEGER[],
    external_revision_subject INTEGER,
    revisor INTEGER[] NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    revision_type INTEGER NOT NULL,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
