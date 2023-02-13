CREATE TABLE users (
  id serial PRIMARY KEY,
  name text NOT NULL,
  pix_key text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE clients (
  id serial PRIMARY KEY,
  name text NOT NULL,
  whatsapp text NOT NULL,
  email text NOT NULL,
  user_id integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE charges (
  id serial PRIMARY KEY,
  value text NOT NULL,
  observation text NOT NULL,
  notification_date timestamp NOT NULL,
  user_id integer NOT NULL,
  client_id integer NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE clients ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE charges ADD FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE;
ALTER TABLE charges ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

CREATE OR REPLACE FUNCTION update_edit_date()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_edit_date_trigger
BEFORE UPDATE ON clients
FOR EACH ROW
EXECUTE FUNCTION update_edit_date();

CREATE TRIGGER update_edit_date_trigger
BEFORE UPDATE ON charges
FOR EACH ROW
EXECUTE FUNCTION update_edit_date();

CREATE TRIGGER update_edit_date_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_edit_date();
