-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE DROP COLUMN last_fetched_at;
-- +goose StatementEnd
