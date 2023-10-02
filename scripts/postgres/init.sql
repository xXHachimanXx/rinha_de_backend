CREATE TABLE
    IF NOT EXISTS public.person (
        id uuid PRIMARY KEY NOT NULL,
        nickname varchar(32) UNIQUE NOT NULL,
        "name" varchar(100) NOT NULL,
        birthdate date NOT NULL,
        stack text
    );