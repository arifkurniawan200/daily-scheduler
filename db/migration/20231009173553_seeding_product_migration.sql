-- +goose Up
-- +goose StatementBegin
INSERT INTO products (name, type, model_product, price, profit, product_manager, release_date)
VALUES
    ('Sucorinvest Sharia Sustainability', 'Saham', 'Konvensional', 100.00, 0.14, 'Succor Asset Management', '2023-10-10 00:00:00'),
    ('TRIM Kapital Plus', 'Pasar Uang', 'Syariah', 150.00, 0.11, 'Succor Asset Management', '2023-10-10 00:00:00'),
    ('Trimegah Balanced Absolute Strategy Kelas A', 'Pendapatan Tetap', 'Konvensional', 200.00, 0.19, 'Succor Asset Management', '2023-10-20 00:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
