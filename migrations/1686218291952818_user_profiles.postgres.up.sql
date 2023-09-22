DROP TABLE if exists user_profiles;

CREATE TABLE user_profiles (
    id serial PRIMARY KEY,
    user_account_id int NOT NULL,
    first_name text NOT NULL,
    middle_name text,
    last_name text NOT NULL,
    birth_last_name text,
    father_name text NOT NULL,
    mother_name text NOT NULL,
    mother_birth_last_name text,
    date_of_birth date NOT NULL,
    country_of_birth text NOT NULL,
    city_of_birth text NOT NULL,
    nationality text,
    national_minority text,
    citizenship text NOT NULL,
    address text NOT NULL,
    bank_account text,
    bank_name text,
    official_personal_id text NOT NULL,
    official_personal_document_number text NOT NULL,
    official_personal_document_issuer text NOT NULL,
    gender text NOT NULL,
    single_parent boolean NOT NULL,
    housing_done boolean NOT NULL,
    housing_description text NOT NULL,
    martial_status text NOT NULL,
    date_of_taking_oath date,
    date_of_becoming_judge text,
    active_contract boolean,
    revisor_role boolean,
    engagement_type_id integer,
    is_judge bool,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);
