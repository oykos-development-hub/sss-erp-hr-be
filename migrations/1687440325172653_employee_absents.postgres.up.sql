CREATE TABLE employee_absents (
    id serial PRIMARY KEY,
    absent_type_id INT NOT NULL,
    user_profile_id INT NOT NULL,
    target_organization_unit_id INT,
    date_of_start DATE NOT NULL,
    date_of_end DATE NOT NULL,
    description text,
    location text,
    file_id INT,    
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (absent_type_id) REFERENCES absent_types (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (target_organization_unit_id) REFERENCES organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
);
