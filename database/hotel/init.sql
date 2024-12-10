-- Таблица Hotels: хранит информацию об отелях.
CREATE TABLE IF NOT EXISTS hotels (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор отеля

    name TEXT NOT NULL DEFAULT 'Metropol', -- Название отеля

    city TEXT NOT NULL DEFAULT 'Moscow', -- Город нахождения отеля

    hotelier_id INT NOT NULL, -- Идентификатор владельца отеля

    rating INT NOT NULL DEFAULT 5, -- Рейтинг отеля (количество звезд)

    country TEXT NOT NULL DEFAULT 'Russia', -- Страна нахождения отеля

    address TEXT NOT NULL DEFAULT 'Teatral''nyy Proyezd, 2, 109012' -- Адрес отеля
);

-- Таблица RoomsCategories: хранит информацию о категориях номеров в отелях.
CREATE TABLE IF NOT EXISTS rooms_categories (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор категории номеров

    name TEXT NOT NULL DEFAULT 'Standard', -- азвание категории номеров

    occupancy INT NOT NULL DEFAULT 1, -- Максимальное количество человек, которые могут разместиться в номере

    price INT NOT NULL DEFAULT 1000 -- Цена номера за 1 день
);

-- Таблица Rooms: хранит информацию о номерах в отелях.
CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор номера

    hotel_id INT NOT NULL REFERENCES hotels(id) ON DELETE CASCADE, -- Идентификатор отеля
    -- Внешний ключ, ссылающийся на id в таблице Hotels.
    -- Гарантирует, что каждая комната связана с существующим отелем.
    -- ON DELETE CASCADE: автоматически удаляет комнаты, если связанный отель удален.

    number INT NOT NULL, -- Порядковый номер номера в отеле

    room_category_id INT NOT NULL REFERENCES rooms_categories(id) ON DELETE CASCADE -- Идентификатор категории номеров
    -- Внешний ключ, ссылающийся на id в таблице RoomsCategories.
    -- Гарантирует, что каждая комната принадлежит существующей категории.
    -- ON DELETE CASCADE: автоматически удаляет комнаты, если связанная категория удалена.
);
