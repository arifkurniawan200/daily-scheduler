-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id INT AUTO_INCREMENT PRIMARY KEY,
   NIK VARCHAR(255) NOT NULL,
   full_name VARCHAR(255) NOT NULL,
   born_place VARCHAR(255),
   born_date DATE,
   email VARCHAR(255) NOT NULL,
   is_admin BOOLEAN DEFAULT false,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
