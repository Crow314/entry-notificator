CREATE TABLE IF NOT EXISTS cards(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    child_id uuid UNIQUE NOT NULL
        REFERENCES children(id),
    token text NOT NULL,
    token_type text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp,
    UNIQUE(token, token_type)
);
