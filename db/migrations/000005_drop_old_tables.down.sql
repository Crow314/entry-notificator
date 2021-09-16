CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name text UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
CREATE TABLE IF NOT EXISTS cards(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid UNIQUE NOT NULL
        REFERENCES users(id),
    token text UNIQUE NOT NULL,
    token_type text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
CREATE TABLE IF NOT EXISTS receivers(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid UNIQUE NOT NULL
        REFERENCES users(id),
    address text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
