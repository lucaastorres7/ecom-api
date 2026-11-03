CREATE TABLE IF NOT EXISTS order_items (
  "id"        SERIAL PRIMARY KEY,
  "orderId"   INT NOT NULL,
  "productId" INT NOT NULL,
  "quantity"  INT NOT NULL,
  "price"     DECIMAL(10,2) NOT NULL,

  CONSTRAINT fk_order_items_order
    FOREIGN KEY ("orderId")
    REFERENCES orders ("id")
    ON DELETE CASCADE,
  
  CONSTRAINT fk_order_items_product
    FOREIGN KEY ("productId")
    REFERENCES products ("id")
);