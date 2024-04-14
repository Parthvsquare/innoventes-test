CREATE TABLE MusicAlbums (
    album_id uuid DEFAULT gen_random_uuid(),
    album_name VARCHAR(255) NOT NULL CHECK (LENGTH(album_name) >= 5),
    release_date DATE NOT NULL,
    genre VARCHAR(100),
    price NUMERIC(6,2) NOT NULL CHECK (price >= 100 AND price <= 1000),
    description TEXT,

  PRIMARY KEY (album_id)
);

CREATE TABLE Musicians (
    musician_id uuid DEFAULT gen_random_uuid(),
    musician_name VARCHAR(255) NOT NULL CHECK (LENGTH(musician_name) >= 3),
    musician_type VARCHAR(100) NOT NULL,

  PRIMARY KEY (musician_id)
);

CREATE TABLE AlbumMusicians (
    album_id uuid REFERENCES MusicAlbums(album_id),
    musician_id uuid REFERENCES Musicians(musician_id),
    PRIMARY KEY (album_id, musician_id)
);

