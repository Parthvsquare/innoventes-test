-- name: GetAlbumsByMusicianId :many
SELECT * FROM album_musician_view WHERE musician_id = $1;

-- name: GetAlbumByAlbumId :one
SELECT * FROM MusicAlbums WHERE album_id = $1;

-- name: AddNewAlbum :one
INSERT INTO MusicAlbums (album_name, release_date, genre, price, description) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateAlbum :one
UPDATE
  MusicAlbums
SET
  album_name = COALESCE($1, album_name),
  release_date = COALESCE($2, release_date),
  genre = COALESCE($3, genre),
  price = COALESCE($4, price),
  description = COALESCE($5, description)
WHERE album_id = $6
RETURNING *;

-- name: DeleteAlbum :exec
DELETE FROM MusicAlbums WHERE album_id = $1;

-- name: AddMusicianToAlbum :exec
INSERT INTO AlbumMusicians (album_id, musician_id) VALUES ($1, $2);

-- name: DeleteMusicianFromAlbum :exec
DELETE FROM AlbumMusicians WHERE album_id = $1 AND musician_id = $2;

-- name: UpdateMusicianOfAlbum :exec
UPDATE AlbumMusicians
SET
 musician_id = COALESCE(@old_musician_id, musician_id)
 WHERE album_id = $1 AND musician_id = @new_musician_id;
