CREATE TABLE IF NOT EXISTS employee_experiences (
    id serial PRIMARY KEY,
    user_profile_id INT NOT NULL,
    relevant boolean NOT NULL DEFAULT false,
    organization_unit TEXT,
    organization_unit_id INT,
    amount_of_experience INT,
    amount_of_insured_experience INT,
    date_of_signature DATE,
    date_of_start DATE,
    date_of_end DATE,
    file_id INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (organization_unit_id) REFERENCES organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
);
