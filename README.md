# Task for a Great Backend Developer - Scootin' Aboot

<hr style="border:2px solid"> 

### 1. Requirements:

- docker engine version >= 23.0.5
- docker composer version >= 2.17.3

### 2. Installation

- git clone https://github.com/boolfixer/scootin_aboot.git && cd scootin_aboot
- cp .env.dist .env
- docker-compose up -d
- docker exec -it scootin_aboot__db bash
- mysql -uroot -proot scootin_aboot < /scootin_aboot.sql && exit

### 3. Run application

- docker exec -it scootin_aboot__go bash 
- cd /app
- go run cmd/main.go

<hr style="border:2px solid"> 

# Application description:

### 1. Base concepts:

- There are 3 tables: `users`, `scooters` and `scooters_occupations`. 
- `scooters_occupations` table contains unique index on `scooter_id` field.
- If there is record in `scooters_occupations`, then such scooter is considered as occupied.
- On `scooter release` request we remove record from `scooters_occupations` table.

### 2. Application structure

1. Application entrypoint placed in `cmd/main.go`
2. database migrations placed in `migration/` folder 
3. Application logic placed in `internal/` folder:
   1. `internal/container` - bootstrap DI container
   2. `internal/controller` - entrypoint to serve HTTP requests
   3. `internal/database` - connection to db
   4. `internal/dto` - objects to represent http request data 
   5. `internal/http_client` - http client for user requests simulation
   6. `internal/http_error` - custom errors to be proceed in error middleware (see the next point)
   7. `internal/middleware` - contains 2 middlewares for user authentication and error handling
   8. `internal/http_server` - [gin](https://github.com/gin-gonic/gin) wrapper
   9. `internal/mock` - mock objects for unit testing
   10. `internal/model` - classes to represent database tables
   11. `internal/repository` - responsible for communication with database
   12. `internal/repository` - business logic
   13. `internal/simulator` - user behaviour simulation

### 3. Used vendors / open-source solutions:

1. [uber-go](https://github.com/uber-go)
2. [gin](https://github.com/gin-gonic/gin) 
3. [database/sql](https://pkg.go.dev/database/sql)

<hr style="border:2px solid"> 

# Potential places to be improved:
1. Use ORM like [gorm](https://gorm.io/index.html) for entity management and sql injection prevention
2. Add intergenerational and functional tests.
3. Add logging for more descriptive error investigation.
4. Error handling with more golang standards.
