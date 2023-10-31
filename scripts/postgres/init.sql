CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE OR REPLACE FUNCTION generate_searchable(_nome VARCHAR, _apelido VARCHAR, _stack JSON)
    RETURNS TEXT AS $$
    BEGIN
    RETURN _nome || _apelido || _stack;
    END;
$$ LANGUAGE plpgsql IMMUTABLE;

CREATE TABLE
    IF NOT EXISTS person (
        id uuid PRIMARY KEY NOT NULL,
        nickname varchar(32) UNIQUE NOT NULL,
        "name" varchar(100) NOT NULL,
        birthdate date NOT NULL,
        stack JSON,
        searchable text GENERATED ALWAYS AS (generate_searchable(name, nickname, stack)) STORED
    );

CREATE INDEX IF NOT EXISTS idx_pessoas_searchable ON person USING gist (searchable public.gist_trgm_ops (siglen='64'));

CREATE UNIQUE INDEX IF NOT EXISTS idx_pessoas_apelido ON person USING btree (nickname);

