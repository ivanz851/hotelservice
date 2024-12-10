-- Таблица Bookings: хранит информацию о бронированиях.
CREATE TABLE bookings  (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор бронирования
    client_id INT NOT NULL DEFAULT 0, -- Идентификатор клиента, оформившего бронирование
    hotel_id INT NOT NULL, -- Идентификатор отеля, в котором бронируется номер
    room_id INT NOT NULL, -- Идентификатор бронируемого номера
    booking_start_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(), -- Временная метка заселения клиента
    booking_finish_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() + INTERVAL '1 day', -- Временная метка выселения клиента
    creation_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() -- Временная метка бронирования
);
