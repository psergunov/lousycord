CREATE TABLE IF NOT EXISTS users (
    id               BIGSERIAL       PRIMARY KEY,
    username         VARCHAR(255)    NOT NULL,
    display_name     VARCHAR(255)    NOT NULL,
    password         VARCHAR(255)    NOT NULL,
);