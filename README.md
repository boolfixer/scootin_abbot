# Task for a Great Backend Developer - Scootin' Aboot

<hr style="border:2px solid"> 

### 1. Requirements:

- docker engine version >= 23.0.5
- docker composer version >= 2.17.3

### 2. Installation

- git clone https://github.com/boolfixer/scootin_aboot.git && cd scootin_aboot
- cp .env.dist .env
- docker-compose up -d
- docker exec scootin_aboot__db bash
- mysql -uroot -proot scootin_aboot < /scootin_aboot.sql

### 3. Run application

- docker exec -it scootin_aboot__go bash 
- cd /app
- go run cmd/main.go

<hr style="border:2px solid"> 

# Application description:

### 1. Application structure

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

### 2. Used vendors / open-source solutions:

1. [uber-go](https://github.com/uber-go)
2. [gin](https://github.com/gin-gonic/gin) 
3. [database/sql](https://pkg.go.dev/database/sql)
