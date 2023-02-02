-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

-- name: GetLacuba :one
SELECT * FROM lacuba
where id = ? LIMIT 1;

-- name: CreateLacuba :execresult
INSERT INTO lacuba (
    longtitude,
    latitude

) VALUES (
             ?, ?
         );

-- name: ListLacubas :many
SELECT * FROM lacuba;