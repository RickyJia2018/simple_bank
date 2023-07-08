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

## Database preparation

Get Postgres Image： 

`docker pull postgres` 

Run Postgres container ：
`docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=guessMyPW -d postgres:12-alpine`