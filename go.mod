module github.com/chouandy/dbmigrater

go 1.16

require (
	github.com/chouandy/go-sdk v1.1.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/lib/pq v1.10.2
	github.com/mitchellh/cli v1.1.2
)

// replace github.com/chouandy/go-sdk => ../go-sdk
