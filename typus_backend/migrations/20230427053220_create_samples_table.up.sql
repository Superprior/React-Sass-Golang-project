CREATE TABLE code_samples (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(1000) ARRAY NOT NULL,
    language VARCHAR(255) NOT NULL
);