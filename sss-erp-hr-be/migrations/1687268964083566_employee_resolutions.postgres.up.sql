CREATE TABLE IF NOT EXISTS employee_resolutions (
    id serial PRIMARY KEY,
    resolution_type_id INT NOT NULL,
    user_profile_id INT NOT NULL,
    resolution_purpose TEXT,
    date_of_start DATE NOT NULL,
    date_of_end DATE NOT NULL,
    file_id INT,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE
);
