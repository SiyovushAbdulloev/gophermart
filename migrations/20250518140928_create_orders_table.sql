-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    id bigint PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    points NUMERIC(10, 2) DEFAULT 0,
    status varchar(255) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
