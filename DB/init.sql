-- ============================================
-- МИГРАЦИЯ: Создание таблиц театра
-- ============================================

-- 1. Создание ENUM типа
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'booking_status' AND typnamespace = (SELECT oid FROM pg_namespace WHERE nspname = 'theatre')) THEN
        CREATE TYPE theatre.booking_status AS ENUM ('AVAILABLE', 'RESERVED', 'BOOKED', 'UNAVAILABLE');
    END IF;
END $$;

-- 2. Театры
CREATE TABLE IF NOT EXISTS theatre.theatres (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

-- Column comments

COMMENT ON COLUMN theatre.theatres.id IS 'Идентификатор';
COMMENT ON COLUMN theatre.theatres."name" IS 'Название';

-- 3. Спектакли
CREATE TABLE IF NOT EXISTS theatre.spectacles (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description TEXT,
    genre VARCHAR,
    age_limit INT,
    duration_minutes INT,
    pushkin_card BOOLEAN DEFAULT FALSE,
    theatre_id INT,
    preview_url VARCHAR,
    CONSTRAINT spectacles_theatres_fk FOREIGN KEY (theatre_id) REFERENCES theatre.theatres(id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- Column comments

COMMENT ON COLUMN theatre.spectacles.id IS 'Идентификатор';
COMMENT ON COLUMN theatre.spectacles.description IS 'Описание';
COMMENT ON COLUMN theatre.spectacles."name" IS 'Название';
COMMENT ON COLUMN theatre.spectacles.genre IS 'Жанр';
COMMENT ON COLUMN theatre.spectacles.age_limit IS 'Возрастное ограничение';
COMMENT ON COLUMN theatre.spectacles.duration_minutes IS 'Длительность в минутах';
COMMENT ON COLUMN theatre.spectacles.pushkin_card IS 'Флаг пушкинской карты';
COMMENT ON COLUMN theatre.spectacles.theatre_id IS 'Идентификатор театра';
COMMENT ON COLUMN theatre.spectacles.preview_url IS 'Ссылка на превью';

-- 4. Актеры
CREATE TABLE IF NOT EXISTS theatre.actors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    middle_name VARCHAR,
    last_name VARCHAR NOT NULL
);

-- Column comments

COMMENT ON COLUMN theatre.actors.id IS 'идентификатор';
COMMENT ON COLUMN theatre.actors.first_name IS 'имя актёра';
COMMENT ON COLUMN theatre.actors.middle_name IS 'отчество актёра';
COMMENT ON COLUMN theatre.actors.last_name IS 'фамилия актёра';

-- 5. Актеры-спектакли
CREATE TABLE IF NOT EXISTS theatre.actors_spectacle (
    actor_id INT NOT NULL,
    spectacle_id INT NOT NULL,
    role VARCHAR,
    CONSTRAINT actors_spectacle_pk PRIMARY KEY (actor_id, spectacle_id),
    CONSTRAINT actors_spectacle_actors_fk FOREIGN KEY (actor_id) REFERENCES theatre.actors(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT actors_spectacle_spectacles_fk FOREIGN KEY (spectacle_id) REFERENCES theatre.spectacles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Column comments

COMMENT ON COLUMN theatre.actors_spectacle.actor_id IS 'Идентификатор актёра';
COMMENT ON COLUMN theatre.actors_spectacle.spectacle_id IS 'Идентификатор спектакля';
COMMENT ON COLUMN theatre.actors_spectacle."role" IS 'Роль в спектакле';

-- 6. Площадки
CREATE TABLE IF NOT EXISTS theatre.platforms (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    address VARCHAR,
    scheme TEXT
);

-- Column comments

COMMENT ON COLUMN theatre.platforms.id IS 'Идентификатор';
COMMENT ON COLUMN theatre.platforms.scheme IS 'Схема посадки';
COMMENT ON COLUMN theatre.platforms.address IS 'Адрес';
COMMENT ON COLUMN theatre.platforms."name" IS 'Название';

-- 7. Группы мест
CREATE TABLE IF NOT EXISTS theatre.seats_groups (
    id SERIAL PRIMARY KEY,
    color VARCHAR,
    price NUMERIC
);

-- Column comments

COMMENT ON COLUMN theatre.seats_groups.id IS 'Идентификатор';
COMMENT ON COLUMN theatre.seats_groups.color IS 'Цвет';
COMMENT ON COLUMN theatre.seats_groups.price IS 'Цена';

-- 8. Показы
CREATE TABLE IF NOT EXISTS theatre.shows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    spectacle_id INT,
    platform_id INT,
    date TIMESTAMP,
    CONSTRAINT show_spectacles_fk FOREIGN KEY (spectacle_id) REFERENCES theatre.spectacles(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT show_platforms_fk FOREIGN KEY (platform_id) REFERENCES theatre.platforms(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Column comments

COMMENT ON COLUMN theatre.shows.id IS 'Иднтификатор';
COMMENT ON COLUMN theatre.shows.spectacle_id IS 'Идентификатор спектакля';
COMMENT ON COLUMN theatre.shows."date" IS 'Дата';
COMMENT ON COLUMN theatre.shows.platform_id IS 'Идентификатор площадки';

-- 9. Места
CREATE TABLE IF NOT EXISTS theatre.seats (
    id SERIAL PRIMARY KEY,
    number INT,
    row INT,
    type VARCHAR,
    group_id INT,
    show_id UUID,
    status theatre.booking_status DEFAULT 'AVAILABLE',
    CONSTRAINT seats_seats_groups_fk FOREIGN KEY (group_id) REFERENCES theatre.seats_groups(id) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT seats_show_fk FOREIGN KEY (show_id) REFERENCES theatre.shows(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Column comments

COMMENT ON COLUMN theatre.seats.id IS 'Идентификатор';
COMMENT ON COLUMN theatre.seats."number" IS 'Номер';
COMMENT ON COLUMN theatre.seats."row" IS 'Ряд';
COMMENT ON COLUMN theatre.seats."type" IS 'Тип';
COMMENT ON COLUMN theatre.seats.group_id IS 'Идентификатор группы мест';
COMMENT ON COLUMN theatre.seats.show_id IS 'Идентификатор показа';
COMMENT ON COLUMN theatre.seats.status IS 'Статус бронирования';

-- ============================================
-- ИНДЕКСЫ
-- ============================================

CREATE INDEX IF NOT EXISTS idx_seats_show_id ON theatre.seats(show_id);
CREATE INDEX IF NOT EXISTS idx_seats_status ON theatre.seats(status);
CREATE INDEX IF NOT EXISTS idx_shows_date ON theatre.shows(date);
CREATE INDEX IF NOT EXISTS idx_shows_spectacle_id ON theatre.shows(spectacle_id);
CREATE INDEX IF NOT EXISTS idx_spectacles_theatre_id ON theatre.spectacles(theatre_id);

-- ============================================
-- КОММЕНТАРИИ
-- ============================================

COMMENT ON TABLE theatre.theatres IS 'Театры';
COMMENT ON TABLE theatre.spectacles IS 'Спектакли';
COMMENT ON TABLE theatre.actors IS 'Актеры';
COMMENT ON TABLE theatre.actors_spectacle IS 'Связь актеров и спектаклей';
COMMENT ON TABLE theatre.platforms IS 'Площадки/залы';
COMMENT ON TABLE theatre.seats_groups IS 'Группы мест (ценовые категории)';
COMMENT ON TABLE theatre.shows IS 'Показы спектаклей';
COMMENT ON TABLE theatre.seats IS 'Места в зале на конкретный показ';