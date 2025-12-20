-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "order" RENAME amount to price;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "order" RENAME price to amount;
-- +goose StatementEnd
