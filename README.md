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
```
