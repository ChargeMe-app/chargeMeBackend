CREATE TABLE amenities
(
    id          TEXT PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES places (id) ON DELETE CASCADE,
    type        INTEGER NOT NULL,
    created_at  TIMESTAMP(0) NOT NULL
);