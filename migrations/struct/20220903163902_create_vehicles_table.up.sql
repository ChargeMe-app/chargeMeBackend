CREATE TABLE vehicles
(
    user_id      TEXT REFERENCES users (id) ON DELETE CASCADE,
    vehicle_type TEXT NOT NULL
);