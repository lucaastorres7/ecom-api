CREATE TYPE order_status AS ENUM (
  'pending',
  'completed',
  'cancelled'
);

CREATE TABLE IF NOT EXISTS orders (
  "id" SERIAL PRIMARY KEY,
  "userId" INT NOT NULL,
  "total" INT NOT NULL,
  "status" order_status NOT NULL DEFAULT 'pending',
  "address" TEXT NOT NULL,
  "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_orders_user
    FOREIGN KEY ("userId")
    REFERENCES users ("id")
    ON DELETE CASCADE
);