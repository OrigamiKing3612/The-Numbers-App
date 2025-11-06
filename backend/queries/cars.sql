-- name: GetAllCars :many
select * from cars;

-- name: UpdateCheckedIn :exec
update cars set checked_in = ? where id = ?;

