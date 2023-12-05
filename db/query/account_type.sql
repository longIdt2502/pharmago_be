-- name: GetAccountType :one
SELECT * from account_type
WHERE id = $1 LIMIT 1;