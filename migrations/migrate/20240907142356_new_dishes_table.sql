-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dishes
(
    id          SERIAL PRIMARY KEY,     -- Уникальный идентификатор для каждого блюда
    name        VARCHAR(100)  NOT NULL, -- Название блюда
    category    VARCHAR(50)   NOT NULL, -- Категория блюда (например, 'Закуски', 'Основные блюда', 'Десерты')
    price       DECIMAL(5, 2) NOT NULL, -- Цена блюда
    description TEXT                    -- Описание блюда
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dishes;
-- +goose StatementEnd
