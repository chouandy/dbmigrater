module github.com/chouandy/dbmigrater

go 1.12

require (
	github.com/chouandy/go-sdk v0.0.0-20191226140241-c7b9130fc762
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang-migrate/migrate/v4 v4.7.1
	github.com/mattn/go-colorable v0.1.0 // indirect
	github.com/mitchellh/cli v1.0.0
)

// replace github.com/chouandy/go-sdk => ../go-sdk
