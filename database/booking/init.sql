-- Таблица Bookings: хранит информацию о бронированиях.
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор бронирования
    client_id INT NOT NULL DEFAULT 0, -- Идентификатор клиента, оформившего бронирование
    hotel_id INT NOT NULL, -- Идентификатор отеля, в котором бронируется номер
    status TEXT NOT NULL -- Статус бронирования
);
