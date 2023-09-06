CREATE TABLE job_positions (
  id serial PRIMARY KEY,
  title text NOT NULL,
  abbreviation text,
  serial_number text,
  requirements text,
  description text,
  is_judge BOOL NOT NULL DEFAULT false,
  is_judge_president BOOL NOT NULL DEFAULT false,
  color text,
  icon text,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);