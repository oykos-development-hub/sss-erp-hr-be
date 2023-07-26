CREATE TABLE employee_educations (
    id serial PRIMARY KEY,
    user_profile_id INT NOT NULL,
    education_type_id INT NOT NULL,
    date_of_certification DATE,
    price INT,
    date_of_start DATE,
    date_of_end DATE,
    academic_title text,
    expertise_level text,
    certificate_issuer text,
    description text,
    title text,
    file_id int,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE
);
