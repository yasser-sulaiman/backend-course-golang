# Docker 
## compose up
`docker-compose up --build`

## clean volumes
`docker-compose down -v`

# Migrations
## Create
`migrate create -seq -ext sql -dir .\cmd\migrate\migrations\ create_users` 

## Run
### up
`migrate -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost/social?sslmode=disable" up`
### down
`migrate -path cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost/social?sslmode=disable" down`