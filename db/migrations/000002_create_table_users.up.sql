CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name text UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
