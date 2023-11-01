
run sqlc with docker


docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
