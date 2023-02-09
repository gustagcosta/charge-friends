CREATE TABLE clients (
  id serial PRIMARY KEY,
  name text NOT NULL,
  description text DEFAULT NULL,
  whatsapp_number integer DEFAULT NULL,
  email text DEFAULT NULL,
  user_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
  id serial PRIMARY KEY,
  name text NOT NULL,
  corporate_name text NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE products (
  id serial PRIMARY KEY,
  title text NOT NULL,
  description text DEFAULT NULL,
  price integer NOT NULL,
  profit smallint NOT NULL,
  user_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE sales (
  id serial PRIMARY KEY,
  paid_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  client_id integer NOT NULL,
  user_id integer NOT NULL
);

CREATE TABLE sales_products (
  sale_id integer NOT NULL,
  product_id integer NOT NULL
);

ALTER TABLE clients ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE products ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE sales ADD FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE;
ALTER TABLE sales ADD FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE sales_products ADD FOREIGN KEY (sale_id) REFERENCES sales(id) ON DELETE CASCADE;
ALTER TABLE sales_products ADD FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;
