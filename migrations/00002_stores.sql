-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE stores (
        id BIGSERIAL PRIMARY KEY NOT NULL,
        org_id bigint NOT NULL REFERENCES orgs(id),
        name character varying(100) NOT NULL,
        created_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        updated_at timestamptz DEFAULT statement_timestamp() NOT NULL
);
COMMENT ON TABLE stores IS '店舗';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS stores CASCADE;
