CREATE TABLE places
(
    id        TEXT PRIMARY KEY,
    name      TEXT  NOT NULL,
    score     FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    latitude  FLOAT NOT NULL
);

COMMENT ON TABLE places IS 'таблица с локациями зарядных станций';
COMMENT ON COLUMN places.id IS 'идентификатор локации';
COMMENT ON COLUMN places.name IS 'название локации';
COMMENT ON COLUMN places.score IS 'рейтинг локации';
COMMENT ON COLUMN places.longitude IS 'долгота';
COMMENT ON COLUMN places.latitude IS 'широта';