-- Таблица Hotels: хранит информацию об отелях.
CREATE TABLE IF NOT EXISTS hotels (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор отеля
    name TEXT NOT NULL, -- Название отеля
    address TEXT NOT NULL, -- Адрес отеля
    price_per_night DOUBLE PRECISION, -- Цена за ночь
    email TEXT NOT NULL -- Электронная почта отеля
);
