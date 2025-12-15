-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
    CREATE EXTENSION IF NOT EXISTS "pgcrypto";
    CREATE TABLE orders(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id UUID,
        amount DOUBLE PRECISION,
        status TEXT,
        create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
    DROP TABLE IF EXISTS order;

-- +goose StatementEnd
