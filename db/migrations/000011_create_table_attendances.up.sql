CREATE TABLE IF NOT EXISTS attendances(
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    child_id uuid NOT NULL
        REFERENCES children(id),
    place_id uuid NOT NULL
        REFERENCES places(id),
    is_entry boolean NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);
