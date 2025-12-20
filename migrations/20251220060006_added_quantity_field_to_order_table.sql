-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "order" ADD COLUMN quantity int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "order" DROP COLUMN quantity;
-- +goose StatementEnd
