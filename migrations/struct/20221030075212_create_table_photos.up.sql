CREATE TABLE location_photos
(
    id          TEXT PRIMARY KEY,
    user_id     uuid REFERENCES users (id) ON DELETE CASCADE,
    name        TEXT         NOT NULL,
    location_id TEXT REFERENCES places (id) ON DELETE CASCADE,
    caption     TEXT,
    created_at  TIMESTAMP(0) NOT NULL
);