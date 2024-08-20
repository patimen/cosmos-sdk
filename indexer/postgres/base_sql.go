package postgres

// BaseSQL is the base SQL that is always included in the schema.
const BaseSQL = `
CREATE OR REPLACE FUNCTION nanos_to_timestamptz(nanos bigint) RETURNS timestamptz AS $$
    SELECT to_timestamp(nanos / 1000000000) + (nanos / 1000000000) * INTERVAL '1 microsecond'
$$ LANGUAGE SQL IMMUTABLE;

CREATE TABLE IF NOT EXISTS block
(
    number BIGINT NOT NULL PRIMARY KEY,
    header_json JSONB  NULL
	header_bytes BYTEA NULL
);

CREATE TABLE IF NOT EXISTS tx
(
    id               BIGSERIAL PRIMARY KEY,
    block_number     BIGINT NOT NULL REFERENCES block (number),
    index_in_block   BIGINT NOT NULL,
    data_json        JSONB  NOT NULL
	data_bytes       BYTEA  NOT NULL
);

CREATE TABLE IF NOT EXISTS event
(
    id           BIGSERIAL PRIMARY KEY,
    block_number BIGINT NOT NULL REFERENCES block (number),
    tx_id        BIGINT NULL REFERENCES tx (id),
    msg_index    BIGINT NULL,
    event_index  BIGINT NULL,
    type         TEXT   NOT NULL,
    data         JSONB  NOT NULL
);
`
