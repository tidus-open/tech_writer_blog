module tdao

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.1 // indirect
	tutil v0.0.0-00010101000000-000000000000 // indirect
)

replace tutil => ../util
