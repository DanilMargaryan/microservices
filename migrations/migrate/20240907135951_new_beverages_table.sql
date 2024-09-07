-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS beverages
(
    id          SERIAL PRIMARY KEY,     -- Уникальный идентификатор для каждого напитка
    name        VARCHAR(100)  NOT NULL, -- Название напитка
    type        VARCHAR(50)   NOT NULL, -- Тип напитка: 'Чай' или 'Кофе'
    price       DECIMAL(5, 2) NOT NULL, -- Цена напитка (например, 5.99)
    description TEXT                    -- Описание напитка
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS beverages
-- +goose StatementEnd
