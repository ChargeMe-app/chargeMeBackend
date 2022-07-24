CREATE TABLE stations(
    id TEXT PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES places (id) ON DELETE CASCADE
);

COMMENT ON TABLE stations IS 'таблица с зарядными станциями';
COMMENT ON COLUMN stations.id IS 'идентификатор станции';
COMMENT ON COLUMN stations.location_id IS 'локация, в которой находиться станция';