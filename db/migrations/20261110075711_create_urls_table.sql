-- +goose Up
-- +goose StatementBegin
CREATE TABLE  urls (
    id         BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(16) NOT NULL UNIQUE,
    long_url   TEXT        NOT NULL,
    count      BIGINT      NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ          DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
