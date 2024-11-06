-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE if not exists events
(
    id                     UUID PRIMARY KEY,
    header                 VARCHAR(255) NOT NULL,
    creation_date          TIMESTAMP    NOT NULL,
    start_date             TIMESTAMP    NOT NULL,
    end_date               TIMESTAMP    NOT NULL,
    description            TEXT,
    user_id                VARCHAR(255) NOT NULL,
    delayed_time           INTEGER,
    notification_send_date TIMESTAMP
    );
-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists events;
SELECT 'down SQL query';
-- +goose StatementEnd
