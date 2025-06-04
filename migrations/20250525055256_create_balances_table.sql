-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS balances (
    user_id bigint not null,
    amount NUMERIC(10, 2) default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS balances;
-- +goose StatementEnd
