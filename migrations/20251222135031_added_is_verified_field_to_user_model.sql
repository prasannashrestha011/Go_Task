-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "user" ADD COLUMN is_verified boolean DEFAULT false
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "user" DROP COLUMN is_verified
-- +goose StatementEnd
