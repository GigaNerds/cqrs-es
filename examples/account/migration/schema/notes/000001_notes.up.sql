--
-- Create tables of accounts domain.
--

--
-- Schema for `accounts` table.
--

CREATE TABLE IF NOT EXISTS accounts (
    id                  UUID NOT NULL PRIMARY KEY,
    title               NAME,
    note_content        TEXT,
    created_at          TIMESTAMPTZ NOT NULL,
    deleted_at          TIMESTAMPTZ
);

CREATE INDEX deleted_at_idx ON accounts (deleted_at);

--
-- Schema for table `accounts_events`.
--

CREATE TABLE IF NOT EXISTS accounts_events (
    id           UUID NOT NULL PRIMARY KEY
                 CHECK (id > '00000000-0000-0000-0000-000000000000'),
    type         TEXT NOT NULL,
    account_id   UUID NOT NULL,
    data         JSONB NOT NULL,
    happened_at  TIMESTAMPTZ NOT NULL
);

CREATE INDEX note_id_idx ON accounts_events (account_id);
