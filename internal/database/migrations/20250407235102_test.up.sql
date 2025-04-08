CREATE TABLE animes (
    uuid UUID DEFAULT gen_random_uuid(),
    series_id VARCHAR(255) NOT NULL,
    series_title VARCHAR(255) NOT NULL,
    episodes INT NOT NULL,
    created_at TIMESTAMP(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (uuid)
);
