CREATE TABLE animes (
    uuid UUID DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    season_number INT NOT NULL,
    crunchyroll_series_id VARCHAR(255) NOT NULL,
    episodes INT NOT NULL,
    created_at TIMESTAMP(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (uuid)
);
