# Task for a Great Backend Developer - Scootin' Aboot

### 1. Requirements:

- docker engine version >= 23.0.5
- docker composer version >= 2.17.3

### 2. Installation

- cp .env.dist .env
- docker-compose up -d
- docker exec scootin_aboot__db bash
- mysql -uroot -proot scootin_aboot < /scootin_aboot.sql

### 3. Run application

- docker exec -it scootin_aboot__go bash 
- cd /app
- go run cmd/main.go