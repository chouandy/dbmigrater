# DB Migrater

DB Migrater is a tool for mysql database migrations.

## Install

```sh
go get -u github.com/chouandy/dbmigrater
```

## DB Config

The db config can be:

- ENV["DB_DRIVER"]
- ENV["DB_HOST"]
- ENV["DB_PORT"]
- ENV["DB_DATABASE"]
- ENV["DB_USERNAME"]
- ENV["DB_PASSWORD"]

## Usage

- Create database

```sh
dbmigrater create
```

- Drop database

```sh
dbmigrater drop
```

- Create a set of timestamped up/down migrations

```sh
dbmigrater migrate create --name {name}
```

- Apply all or N up migrations

```sh
dbmigrater migrate up -n {N}
```

- Apply 1 or N down migrations

```sh
dbmigrater migrate down -n {N}
```

- Drop everything inside database

```sh
dbmigrater migrate drop
```
