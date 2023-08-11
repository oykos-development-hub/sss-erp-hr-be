CREATE TABLE IF NOT EXISTS foreigners (
    id serial PRIMARY KEY,
    user_profile_id INTEGER NOT NULL REFERENCES user_profiles(id),
    work_permit_number TEXT NOT NULL,
    work_permit_issuer TEXT NOT NULL,
    work_permit_date_of_start DATE NOT NULL,
    work_permit_date_of_end DATE,
    work_permit_indefinite_length BOOLEAN NOT NULL,
    residence_permit_date_of_end DATE,
    residence_permit_indefinite_length BOOLEAN NOT NULL,
    residence_permit_number TEXT NOT NULL,
    country_of_origin TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    work_permit_file_id INTEGER,
    residence_permit_file_id INTEGER
);