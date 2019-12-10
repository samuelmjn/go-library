-- DDL generated by Postico 1.5.8
-- Not all database features are supported. Do not use for backup.

-- Table Definition ----------------------------------------------

CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title text,
    author text,
    publisher text,
    is_issued boolean,
    issue_count bigint,
    created_at timestamp without time zone,
    deleted_at timestamp without time zone
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX books_pkey ON books(id int8_ops);
