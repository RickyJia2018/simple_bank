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

Get golang migrate：  `brew install golang-migrate`

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





	



## gRPC
1. install [protobuf compiler & go plugins](https://grpc.io/docs/languages/go/quickstart/)
2. vscode install proto3 plugin
3. Update proto3 setting with 

	```js
	"protoc":{
		"options":[ "--proto_path=proto",]
	}
	```
4. create .proto files in proto
5. generate proto files

	```
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
    ```
6. run go mod tidy
7. brew install evans
8. run `evans --host localhost --port 9090 -r repl`
	
	```
	show service
	call my-RPC-api
	```
    
 
 