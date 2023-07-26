CREATE TABLE employees_in_organization_units (
    id serial PRIMARY KEY,
    user_account_id INT NOT NULL,
    user_profile_id INT NOT NULL,
    position_in_organization_unit_id INT NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (position_in_organization_unit_id) REFERENCES job_positions_in_organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
);
