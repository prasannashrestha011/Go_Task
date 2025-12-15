-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
    ALTER TABLE users RENAME to "user";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
    ALTER TABLE "user" to users;
-- +goose StatementEnd
