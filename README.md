# An simple bank server

## Tech stack

<style>
    .two-column-list {
        column-count: 2;
        column-gap: 20px;
    }
</style>

<div class="two-column-list">
    <ul>
        <li>Postgres</li>
        <li>SQLc</li>
        <li>Gin</li>
        <li>RestAPI</li>
        <li>JWT</li>
    </ul>
    <ul>
        <li>Docker</li>
        <li>K8s</li>
        <li>CI/CD</li>
        <li>GitHub Action</li>
        <li>TDD</li>
    </ul>
</div>
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





	



