# An simple bank server

### Tech stack
* Postgres 
* SQLc
* Gin
* RestAPI
* JWT
* Docker
* K8s

## Database Design
![DBdesign](https://i.imgur.com/HtPnnm6.png)

## Database & Migration

Get Postgres Image： `docker pull postgres` 

Get golang migrate：  `brew install golang-migrate`

Get sqlc: ` brew install sqlc` [Config site](https://docs.sqlc.dev/en/latest/reference/config.html)

Create migration files：

`migrate create -ext sql -dir db/migraiton -seq init_schema`

Run Makefile scripts：

`make postgres`
`make createdb`
`make migrateup`

	



