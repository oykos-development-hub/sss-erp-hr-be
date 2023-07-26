CREATE TABLE job_positions (
  id serial PRIMARY KEY,
  title text NOT NULL,
  abbreviation text NOT NULL,
  serial_number text NOT NULL,
  requirements text NOT NULL,
  description text,
  is_judge BOOL NOT NULL DEFAULT false,
  is_judge_president BOOL NOT NULL DEFAULT false,
  color text,
  icon text,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);