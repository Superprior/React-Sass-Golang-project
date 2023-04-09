CREATE TABLE code_samples (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    lang_id BIGINT NOT NULL,
    CONSTRAINT fk_lang FOREIGN KEY(lang_id) REFERENCES programming_languages(id)
);