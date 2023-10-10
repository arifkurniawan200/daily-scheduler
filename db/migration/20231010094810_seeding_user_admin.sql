-- +goose Up
-- +goose StatementBegin
INSERT INTO users (NIK, full_name, born_place, born_date, email, is_admin, password) VALUES
 ('WnIUt10aFKBZMg==', 'Arif kurniawan', 'Boyolali', '1998-10-10', 'arifkurniawandev96@gmail.com', true, '$2a$10$ye9f2ky/gU/2QGrXRMw4VOTlWSwuVASvIDlq2Rc8wruUpQ7xrwJwm');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
