CREATE TABLE stations(
    id TEXT PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES places (id) ON DELETE CASCADE,
    available INTEGER,
    cost INTEGER,
    name TEXT,
    manufacturer TEXT,
    cost_description TEXT,
    hours TEXT,
    kilowatts FLOAT
);

COMMENT ON TABLE stations IS 'таблица с зарядными станциями';
COMMENT ON COLUMN stations.id IS 'идентификатор станции';
COMMENT ON COLUMN stations.location_id IS 'локация, в которой находиться станция';