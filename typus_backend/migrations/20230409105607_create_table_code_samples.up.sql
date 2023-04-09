CREATE TABLE code_samples (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    lang_slug CHAR(20) NOT NULL,
    CONSTRAINT fk_lang FOREIGN KEY(lang_slug) REFERENCES programming_languages(slug)
);