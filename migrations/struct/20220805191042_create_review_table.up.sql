CREATE TABLE reviews(
    id UUID PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES places (id) ON DELETE CASCADE,
    comment TEXT,
    station_id TEXT NOT NULL REFERENCES stations (id) ON DELETE CASCADE,
    outlet_id TEXT NOT NULL REFERENCES outlets (id) ON DELETE CASCADE,
    rating INT,
    vehicle_name TEXT,
    vehicle_type TEXT,
    created_at       TIMESTAMP(0) NOT NULL
);

COMMENT ON TABLE reviews IS 'таблица с отзывами о локации';
COMMENT ON COLUMN reviews.id IS 'идентификатор отзыва';
COMMENT ON COLUMN reviews.location_id IS 'идентификатор локации';
COMMENT ON COLUMN reviews.comment IS 'комметарий';
COMMENT ON COLUMN reviews.station_id IS 'идентификатор станции';
COMMENT ON COLUMN reviews.outlet_id IS 'идентификатор разъема';
COMMENT ON COLUMN reviews.rating IS 'рейтинг';
