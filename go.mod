module github.com/chouandy/dbmigrater

go 1.12

require (
	github.com/chouandy/go-sdk v0.0.0-20200303110227-c44573fe9b9c
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang-migrate/migrate/v4 v4.7.1
	github.com/lib/pq v1.3.0
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mitchellh/cli v1.0.0
)

// replace github.com/chouandy/go-sdk => ../go-sdk
