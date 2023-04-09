CREATE TYPE programming_language AS ENUM ('TypeScript', 'Python', 'Go', 'Rust', 'Assembly');

CREATE TABLE programming_languages (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    slug CHAR(20) NOT NULL,
    title programming_language
);