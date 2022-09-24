CREATE TABLE reviews
(
    id             TEXT PRIMARY KEY,
    user_id        uuid REFERENCES users (id) ON DELETE CASCADE,
    station_id     TEXT,
    outlet_id      TEXT,
    vehicle_type   INTEGER,
    comment        TEXT,
    kilowatts      FLOAT,
    rating         INT,
    user_name      TEXT,
    connector_type INT,
    vehicle_name   TEXT,
    created_at     TIMESTAMP(0) NOT NULL
);

COMMENT
    ON TABLE reviews IS 'таблица с отзывами о локации';
COMMENT
    ON COLUMN reviews.id IS 'идентификатор отзыва';
COMMENT
    ON COLUMN reviews.comment IS 'комметарий';
COMMENT
    ON COLUMN reviews.station_id IS 'идентификатор станции';
COMMENT
    ON COLUMN reviews.outlet_id IS 'идентификатор разъема';
COMMENT
    ON COLUMN reviews.rating IS 'рейтинг';
