-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "order" ALTER COLUMN amount TYPE DOUBLE PRECISION USING amount::DOUBLE PRECISION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "order" ALTER COLUMN amount TYPE INT USING amount::INT;
-- +goose StatementEnd
