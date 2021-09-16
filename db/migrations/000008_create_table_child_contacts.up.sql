CREATE TABLE IF NOT EXISTS child_contacts(
      id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
      child_id uuid UNIQUE NOT NULL
          REFERENCES children(id),
      email text NOT NULL,
      phone_number text NOT NULL,
      created_at timestamp NOT NULL,
      updated_at timestamp NOT NULL,
      deleted_at timestamp
);
