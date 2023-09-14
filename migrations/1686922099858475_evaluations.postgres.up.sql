CREATE TABLE IF NOT EXISTS evaluations (
    id serial PRIMARY KEY,
    user_profile_id INT NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE,
    evaluation_type_id INT NOT NULL,
    score TEXT NOT NULL,
    date_of_evaluation DATE NOT NULL,
    evaluator TEXT NOT NULL,
    is_relevant BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    file_id INT
);
