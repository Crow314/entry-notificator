CREATE TABLE IF NOT EXISTS child_addresses(
   id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
   child_id uuid UNIQUE NOT NULL
       REFERENCES children(id),
   postal_code text NOT NULL,
   prefecture_code int NOT NULL,
   city text NOT NULL,
   street text NOT NULL,
   room text,
   created_at timestamp NOT NULL,
   updated_at timestamp NOT NULL,
   deleted_at timestamp
);
