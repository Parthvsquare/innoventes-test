CREATE TABLE Music (
    music_id uuid DEFAULT gen_random_uuid(),
    music_name VARCHAR(255) NOT NULL CHECK (LENGTH(music_name) >= 5),
    music_price NUMERIC(6,2) NOT NULL CHECK (music_price >= 100 AND music_price <= 1000),
    music_description TEXT,
    album_id uuid,
    musician_id uuid,

    PRIMARY KEY (music_id),
    CONSTRAINT fk_album_id FOREIGN KEY (album_id) REFERENCES MusicAlbums(album_id),
    CONSTRAINT fk_musician_id FOREIGN KEY (musician_id) REFERENCES Musicians(musician_id)
);

-- creating view

CREATE VIEW MusicView AS
SELECT
	mu.music_id,
	mu.music_name,
	mu.music_price,
	mu.music_description,
	ma.album_id,
	ma.album_name,
	ma.release_date,
	ma.genre,
	ma.price,
	ma.description,
	ms.musician_id,
	ms.musician_name,
	ms.musician_type
FROM
	Music mu
	JOIN MusicAlbums ma ON mu.album_id = ma.album_id
	JOIN Musicians ms ON ms.musician_id = mu.musician_id;