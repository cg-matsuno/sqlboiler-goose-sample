-- +goose Up
-- SQL in this section is executed when the migration is applied.

 CREATE TABLE orgs(
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name character varying(50) NOT NULL,
        created_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        updated_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        CONSTRAINT orgs_name_constraint UNIQUE(name)
);

CREATE TABLE users(
       id BIGSERIAL PRIMARY KEY NOT NULL,
       org_id bigint NOT NULL REFERENCES orgs(id),
       uid character varying(100),
       name  character varying(100) NOT NULL,
       owner boolean NOT NULL DEFAULT false,
       created_at timestamptz DEFAULT statement_timestamp() NOT NULL,
       updated_at timestamptz DEFAULT statement_timestamp() NOT NULL,
       CONSTRAINT users_uid_constraint UNIQUE(org_id, uid)
);

 CREATE TABLE groups (
        id BIGSERIAL PRIMARY KEY NOT NULL,
        org_id bigint NOT NULL REFERENCES orgs(id),
        name character varying(100) NOT NULL,
        created_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        updated_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        CONSTRAINT groups_name_constraint UNIQUE(org_id, name)
);
COMMENT ON TABLE groups IS 'ユーザグループ';

 CREATE TABLE groups_users (
        group_id bigint NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        user_id bigint NOT NULL REFERENCES users(id),
        created_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        updated_at timestamptz DEFAULT statement_timestamp() NOT NULL,
        CONSTRAINT groups_users_pkey PRIMARY KEY (group_id, user_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS groups_users;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS orgs;
