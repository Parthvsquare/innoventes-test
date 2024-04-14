-- name: MusiciansById :one
SELECT * FROM Musicians WHERE musician_id = $1;