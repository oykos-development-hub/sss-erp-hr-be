CREATE TABLE IF NOT EXISTS revision_tips (
    id serial PRIMARY KEY,
    revision_id INTEGER NOT NULL REFERENCES revisions(id) ON DELETE CASCADE,
    user_profile_id INTEGER REFERENCES user_profiles(id) ON DELETE CASCADE,
    responsible_person TEXT,
    date_of_accept DATE,
    due_date INTEGER NOT NULL,
    new_due_date INTEGER,
    date_of_reject DATE,
    date_of_execution DATE,
    new_date_of_execution DATE,
    recommendation TEXT NOT NULL,
    status TEXT,
    documents TEXT,
    reasons_for_non_executing TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
