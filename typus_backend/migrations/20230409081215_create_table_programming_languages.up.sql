CREATE TYPE programming_language AS ENUM ('TypeScript', 'Python', 'Go', 'Rust', 'Assembly');

CREATE TABLE programming_languages (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    slug VARCHAR(20) NOT NULL UNIQUE,
    title programming_language
);