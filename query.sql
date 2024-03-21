-- name: GetWord :one
SELECT * FROM words
WHERE id = ? LIMIT 1;

-- name: ListWords :many
SELECT * FROM words
ORDER BY id LIMIT ? OFFSET ?;

-- name: GetWordsRandomly :many
SELECT * FROM words
ORDER BY RANDOM() LIMIT ?;