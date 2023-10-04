CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE OR REPLACE FUNCTION generate_searchable(_nome VARCHAR, _apelido VARCHAR, _stack TEXT[])
    RETURNS TEXT AS $$
    BEGIN
    RETURN _nome || _apelido || _stack;
    END;
$$ LANGUAGE plpgsql IMMUTABLE;

CREATE TABLE
    IF NOT EXISTS people (
        id uuid PRIMARY KEY NOT NULL,
        nickname varchar(32) UNIQUE NOT NULL,
        "name" varchar(100) NOT NULL,
        birthdate date NOT NULL,
        stack text[],
        searchable text GENERATED ALWAYS AS (generate_searchable(name, nickname, stack)) STORED
    );

CREATE INDEX IF NOT EXISTS idx_pessoas_searchable ON people USING gist (searchable public.gist_trgm_ops (siglen='64'));

CREATE UNIQUE INDEX IF NOT EXISTS idx_pessoas_apelido ON people USING btree (nickname);