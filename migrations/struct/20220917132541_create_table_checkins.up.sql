CREATE TABLE checkins
(
    id           UUID PRIMARY KEY NOT NULL,
    user_id      UUID             NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    station_id   TEXT             NOT NULL REFERENCES stations (id) ON DELETE CASCADE,
    outlet_id    TEXT             NOT NULL REFERENCES outlets (id) ON DELETE CASCADE,
    vehicle_type INTEGER,
    comment      TEXT,
    kilowatts    FLOAT,
    rating       INTEGER,
    is_auto       BOOL DEFAULT FALSE,
    user_name    TEXT             NOT NULL,
    duration     INT              NOT NULL,
    started_at   TIMESTAMP(0)     NOT NULL,
    finished_at  TIMESTAMP(0)     NOT NULL,
    deleted_at   TIMESTAMP(0)
);