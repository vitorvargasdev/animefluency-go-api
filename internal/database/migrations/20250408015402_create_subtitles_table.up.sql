CREATE TABLE subtitles (
    uuid UUID DEFAULT gen_random_uuid(),
    anime_uuid UUID NOT NULL,
    episode INT NOT NULL,
    language_code VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (uuid)
);

