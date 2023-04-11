CREATE TABLE code_samples (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content varchar(1000) ARRAY NOT NULL,
    lang_slug VARCHAR(20) NOT NULL,
    CONSTRAINT fk_lang FOREIGN KEY(lang_slug) REFERENCES programming_languages(slug)
);