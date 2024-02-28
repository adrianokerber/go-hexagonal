# go-hexagonal
Project using the Hexagonal Architecture

## Useful commands
```bash
# Run tests
go test ./application/

# Install a specific dependency when not auto imported
go get <module_name>

# Generate mocks for our models
mockgen -destination=application/mocks/application.go -source=application/product.go application

# Access our DB stored on sqlite.db file. Warning: use .quit in order to exit the sqlite shell
sqlite3 sqlite.db
# SQL commands:
create table products(id varchar(255), name varchar(255), price float, status string);
```
