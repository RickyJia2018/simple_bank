# An simple bank server

## Tech stack

* Postgres - SQLc
* RestAPI - Gin - JWT
* TDD
* CI/CD - Github Action
* Docker & K8s


## Database Design
![DBdesign](https://i.imgur.com/HtPnnm6.png)

<details>

<summary>
<h2>Database & Migration</h2>

</summary>

Get Postgres Image： `docker pull postgres` 

Get golang migrate：  `brew install golang-migrate` [Config site](https://github.com/golang-migrate/migrate.git)

Get sqlc: ` brew install sqlc` [Config site](https://docs.sqlc.dev/en/latest/reference/config.html)

Create migration files：

`migrate create -ext sql -dir db/migraiton -seq init_schema`

Run Makefile scripts：

`make postgres` to run postgres databse

`make createdb` to create databse 

`make migrateup` to migrate

`make sqlc` to generate query functions


</details>


## Testing

Run: `make test`

Import: `github.com/stretchr/testify/require` to use require.*




	



