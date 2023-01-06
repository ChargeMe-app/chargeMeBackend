CREATE TABLE places
(
    id                             TEXT PRIMARY KEY,
    name                           TEXT         NOT NULL,
    score                          FLOAT,
    longitude                      FLOAT        NOT NULL,
    latitude                       FLOAT        NOT NULL,
    address                        TEXT,
    icon_type                      TEXT,
    description                    TEXT,
    phone_number                   TEXT,
    access                         INTEGER,
    cost                           BOOL,
    cost_description               TEXT,
    hours                          TEXT,
    open247                        BOOL DEFAULT TRUE,
    is_open_or_active              BOOL DEFAULT TRUE,
    created_at                     TIMESTAMP(0) NOT NULL
);

COMMENT ON TABLE places IS 'таблица с локациями зарядных станций';
COMMENT ON COLUMN places.id IS 'идентификатор локации';
COMMENT ON COLUMN places.name IS 'название локации';
COMMENT ON COLUMN places.score IS 'рейтинг локации';
COMMENT ON COLUMN places.longitude IS 'долгота';
COMMENT ON COLUMN places.latitude IS 'широта';
COMMENT ON COLUMN places.address IS 'адрес';
COMMENT ON COLUMN places.access IS 'доступность';
COMMENT ON COLUMN places.icon_type IS 'тип на иконку';