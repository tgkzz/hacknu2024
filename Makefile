run:
	go run ./cmd/app

vet:
	go vet ./...

swag-exec:
	swag init -g ./cmd/app/main.go --parseDependency --parseInternal -o ./docs

swag-fmt:
	swag fmt

swag-all: swag-fmt swag-exec