-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS withdraws (
    id bigint PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    order_id bigint NOT NULL,
    points bigint DEFAULT 0,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdraws;
-- +goose StatementEnd