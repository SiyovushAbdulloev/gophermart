-- +goose Up
-- +goose StatementBegin
ALTER TABLE balances ADD CONSTRAINT unique_user_id UNIQUE (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE balances DROP CONSTRAINT unique_user_id;
-- +goose StatementEnd
