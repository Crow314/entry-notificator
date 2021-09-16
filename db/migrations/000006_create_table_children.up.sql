CREATE TABLE IF NOT EXISTS children(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    first_name text NOT NULL,
    family_name text NOT NULL,
    birth_date date NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
