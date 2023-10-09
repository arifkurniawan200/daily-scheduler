-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type ENUM('Saham', 'Pasar Uang', 'Pendapatan Tetap') NOT NULL,
  model_product ENUM('Konvensional', 'Syariah') NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  profit DECIMAL(10, 2) NOT NULL,
  product_manager VARCHAR(255) NOT NULL,
  release_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
