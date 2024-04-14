-- name: GetMusicByAlbumId :many
SELECT * FROM MusicView WHERE album_id = $1;

-- name: GetMusicById :one
SELECT * FROM Music WHERE music_id = $1;


-- name: UpdateMusic :one
UPDATE Music
SET
    music_name = COALESCE($1, music_name),
    music_price = COALESCE($2, music_price),
    music_description = COALESCE($3, music_description),
    album_id = COALESCE($4, album_id),
    musician_id = COALESCE($5, musician_id)
WHERE music_id = $6
RETURNING *;

-- name: AddNewMusic :one
INSERT INTO Music (music_name, music_price, music_description, album_id, musician_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteMusic :exec
DELETE FROM Music WHERE music_id = $1;
