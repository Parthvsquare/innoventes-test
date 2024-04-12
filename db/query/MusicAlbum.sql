-- name: GetMusicAlbumsById :one
SELECT * FROM MusicAlbums WHERE album_id = $1;