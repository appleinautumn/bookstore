CREATE TABLE IF NOT EXISTS order_books (
  "order_id" BIGINT NOT NULL,
  "book_id" BIGINT NOT NULL,
  "quantity" INT NOT NULL,
  PRIMARY KEY (order_id, book_id),
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (book_id) REFERENCES books(id)
);
