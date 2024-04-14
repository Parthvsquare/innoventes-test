-- name: GetMusiciansById :one
SELECT * FROM Musicians WHERE musician_id = $1;

-- name: AddNewMusician :one
INSERT INTO Musicians (musician_name, musician_type) VALUES ($1, $2) RETURNING *;

-- name: UpdateMusician :one
UPDATE
  Musicians
SET
  musician_name = COALESCE($1, musician_name),
  musician_type = COALESCE($2, musician_type)
WHERE musician_id = $3
RETURNING *;

-- name: DeleteMusician :exec
DELETE FROM Musicians WHERE musician_id = $1;

-- name: GetMusicByMusicianId :many
SELECT * FROM album_music_musician_view WHERE musician_id = $1;

-- name: GetMusicianByAlbumId :one
SELECT * FROM album_musician_view WHERE album_id = $1;