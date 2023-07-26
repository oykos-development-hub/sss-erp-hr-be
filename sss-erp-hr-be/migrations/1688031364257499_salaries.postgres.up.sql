CREATE TABLE IF NOT EXISTS salaries (
    id serial PRIMARY KEY,
    user_profile_id INTEGER NOT NULL,
    benefited_track BOOLEAN NOT NULL,
    without_raise BOOLEAN NOT NULL,
    insurance_basis TEXT NOT NULL,
    salary_rank TEXT NOT NULL,
    daily_work_hours TEXT NOT NULL,
    weekly_work_hours TEXT NOT NULL,
    education_rank TEXT NOT NULL,
    education_naming TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    user_resolution_id INTEGER,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_resolution_id) REFERENCES employee_resolutions (id) ON UPDATE CASCADE ON DELETE CASCADE
);


