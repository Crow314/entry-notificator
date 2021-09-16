CREATE TABLE IF NOT EXISTS supporter_addresses(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    supporter_id uuid UNIQUE NOT NULL
        REFERENCES supporters(id),
    postal_code text NOT NULL,
    prefecture_code int NOT NULL,
    city text NOT NULL,
    street text NOT NULL,
    room text,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
