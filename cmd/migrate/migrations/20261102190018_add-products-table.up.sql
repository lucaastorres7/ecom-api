CREATE TABLE IF NOT EXISTS products (
  "id"          SERIAL PRIMARY KEY,
  "name"        VARCHAR(255) NOT NULL,
  "description" TEXT NOT NULL,
  "image"       VARCHAR(255) NOT NULL,
  "price"       DECIMAL(10, 2) NOT NULL,
  "quantity"    INT NOT NULL,
  "createdAt"   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);