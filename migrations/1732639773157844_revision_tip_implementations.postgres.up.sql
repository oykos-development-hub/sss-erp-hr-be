CREATE TABLE IF NOT EXISTS revision_tip_implementations (
    id serial PRIMARY KEY,
    tip_id INTEGER NOT NULL REFERENCES revision_tips(id) ON DELETE CASCADE,
    revisor_id INTEGER REFERENCES user_profiles(id) ON DELETE CASCADE,
    new_due_date INTEGER,
    new_date_of_execution DATE,
    status TEXT NOT NULL,
    documents TEXT,
    reasons_for_non_executing TEXT,
    file_ids INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ALTER TABLE revision_tips
-- DROP COLUMN IF EXISTS new_due_date,
-- DROP COLUMN IF EXISTS new_date_of_execution,
-- DROP COLUMN IF EXISTS documents,
-- DROP COLUMN IF EXISTS reasons_for_non_executing;
