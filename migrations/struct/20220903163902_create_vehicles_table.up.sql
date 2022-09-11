CREATE TABLE vehicles
(
    user_id      uuid REFERENCES users (id) ON DELETE CASCADE,
    vehicle_type TEXT NOT NULL
);