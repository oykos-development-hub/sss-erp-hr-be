CREATE TABLE IF NOT EXISTS user_norm_fulfilments (
    id SERIAL PRIMARY KEY,
    user_profile_id INTEGER NOT NULL,
    topic TEXT NOT NULL,
    title TEXT NOT NULL,
    number_of_norm_decrease INTEGER NOT NULL,
    number_of_items INTEGER NOT NULL,
    number_of_items_solved INTEGER NOT NULL,
    evaluation_id INTEGER NOT NULL,
    date_of_evaluation DATE NOT NULL,
    date_of_evaluation_validity DATE NOT NULL,
    relocation_id INTEGER,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles(id),
    FOREIGN KEY (evaluation_id) REFERENCES evaluations(id),
    FOREIGN KEY (relocation_id) REFERENCES employee_absents(id)
);
