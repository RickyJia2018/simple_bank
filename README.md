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

Get Postgres Image： 

`docker pull postgres` 

Get golang migrate：

`brew install golang-migrate`

Create migration files：

`migrate create -ext sql -dir db/migraiton -seq init_schema`

Run Makefile scripts：

`make postgres`
`make createdb`

Migrate db files：
`migrate -path db/migration -database "postgresql://root:myPW@localhost:5432/simple_bank?sslmode=disable" -verbose up`


	



