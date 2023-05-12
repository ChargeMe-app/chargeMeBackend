CREATE TABLE outlets
(
    id         TEXT PRIMARY KEY,
    station_id TEXT         NOT NULL REFERENCES stations (id) ON DELETE CASCADE,
    connector  INTEGER,
    kilowatts  FLOAT,
    power      INTEGER,
    price      FLOAT,
    price_unit TEXT,
    hide       BOOL DEFAULT FALSE,
    created_at TIMESTAMP(0) NOT NULL
);

COMMENT ON TABLE outlets IS 'таблица с разъемами';
COMMENT ON COLUMN outlets.id IS 'идентификатор разъема';
COMMENT ON COLUMN outlets.station_id IS 'идентификатор станции, которой разъем пренадлежит';
COMMENT ON COLUMN outlets.connector IS 'тип разъема';