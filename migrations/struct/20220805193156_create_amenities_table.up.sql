CREATE TABLE amenities
(
    id          UUID PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES places (id) ON DELETE CASCADE,
    type        INTEGER
);