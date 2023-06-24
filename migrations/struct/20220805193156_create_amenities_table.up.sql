CREATE TABLE amenities
(
    id          TEXT PRIMARY KEY,
    location_id TEXT NOT NULL,
    type        INTEGER NOT NULL,
    created_at  TIMESTAMP(0) NOT NULL
);