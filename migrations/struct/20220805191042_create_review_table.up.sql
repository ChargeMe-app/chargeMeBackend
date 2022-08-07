CREATE TABLE reviews
(
    id           TEXT PRIMARY KEY,
    comment      TEXT,
    station_id   TEXT,
    outlet_id    TEXT,
    rating       INT,
    vehicle_name TEXT,
    vehicle_type TEXT,
    created_at   TIMESTAMP(0) NOT NULL
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
