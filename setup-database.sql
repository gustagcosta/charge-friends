CREATE TABLE users (
  id serial PRIMARY KEY,
  name text NOT NULL,
  pix_key text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE clients (
  id serial PRIMARY KEY,
  name text NOT NULL,
  whatsapp text NOT NULL,
  email text NOT NULL,
  user_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE charges (
  id serial PRIMARY KEY,
  description text NOT NULL,
  paid_at timestamptz DEFAULT NULL,
  client_id integer NOT NULL,
  user_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE clients ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE charges ADD FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE;
ALTER TABLE charges ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
