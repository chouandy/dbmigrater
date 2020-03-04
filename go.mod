module github.com/chouandy/dbmigrater

go 1.12

require (
	github.com/chouandy/go-sdk v0.0.0-20200304175702-19748909853b
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate/v4 v4.7.1
	github.com/lib/pq v1.3.0
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mitchellh/cli v1.0.0
)

// replace github.com/chouandy/go-sdk => ../go-sdk
