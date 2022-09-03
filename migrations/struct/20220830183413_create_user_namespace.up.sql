CREATE TABLE users
(
    id              uuid PRIMARY KEY,
    user_identifier TEXT         NOT NULL UNIQUE,
    display_name    TEXT         NOT NULL,
    email           TEXT,
    photo_url       TEXT,
    sign_type       TEXT         NOT NULL,
    created_at      TIMESTAMP(0) NOT NULL
);

CREATE TABLE apple_users
(
    user_id            TEXT REFERENCES users (id) ON DELETE CASCADE,
    authorization_code TEXT NOT NULL,
    identity_token     TEXT NOT NULL
);

CREATE TABLE google_users
(
    user_id      TEXT REFERENCES users (id) ON DELETE CASCADE,
    id_token     TEXT NOT NULL,
    access_token TEXT NOT NULL
);