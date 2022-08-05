CREATE TABLE places
(
    id                             TEXT PRIMARY KEY,
    name                           TEXT  NOT NULL,
    score                          FLOAT,
    longitude                      FLOAT NOT NULL,
    latitude                       FLOAT NOT NULL,
    address                        TEXT,
    icon_type                      TEXT,
    description                    TEXT,
    phone_number                   TEXT,
    access                         INTEGER,
    access_restriction             TEXT,
    access_restriction_description TEXT,
    cost                           BOOL,
    cost_description               TEXT,
    hours                          TEXT,
    open247                        BOOL DEFAULT TRUE,
    isOpenOrActive                 BOOL DEFAULT TRUE
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