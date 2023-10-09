-- +goose Up
-- +goose StatementBegin
CREATE TABLE campaigns (
   id INT AUTO_INCREMENT PRIMARY KEY,
   code VARCHAR(255) NOT NULL,
   name VARCHAR(255) NOT NULL,
   amount DECIMAL(10, 2) NOT NULL,
   start_date DATETIME NOT NULL,
   end_date DATETIME NOT NULL,
   quota INT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE campaigns;
-- +goose StatementEnd
