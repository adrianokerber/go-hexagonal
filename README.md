# go-hexagonal
Project using the Hexagonal Architecture.
> 💡Note: If you prefer, run go just from the container build by docker-compose. For this run `docker compose up -d --build && docker exec -it appproduct sh`

## Useful commands
```bash
# Run tests
go test ./application/

# Install a specific dependency when not auto imported
go get <module_name>

# Auto configure all dependencies removing unused and adding the missing ones
go mod tidy

# Generate mocks for our models
mockgen -destination=application/mocks/application.go -source=application/product.go application

# Access our DB stored on sqlite.db file. Warning: use .quit in order to exit the sqlite shell
sqlite3 sqlite.db
# Quit sqlite3
.quit
# Show all tables from DB
.tables
# SQL commands:
create table products(id varchar(255), name varchar(255), price float, status string);

# Run the app
go run main.go
go run main.go cli -h # For help about the CLI

###
# Using docker for development

# Run the app
docker compose up -d
# Stop the app
docker compose down
# Enter app shell. PS: appproduct is the name of the container
docker exec -it appproduct sh
# Rebuild image
docker compose up -d --build

###
# Using the CLI adapter
go run main.go cli -a=create -n="Product CLI" -p=25.0
go run main.go cli -a=get --id="44e6c4a4-4ff3-411d-895f-3daae1087ea7"
```
