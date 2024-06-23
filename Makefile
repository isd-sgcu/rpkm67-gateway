docker:
	docker-compose up

server:
	go run cmd/main.go

watch: 
	air
mock-gen:
	mockgen -source ./internal/auth/auth.handler.go -destination ./mocks/auth/auth.handler.go
	mockgen -source ./internal/auth/auth.service.go -destination ./mocks/auth/auth.service.go
	mockgen -source ./internal/baan/baan.handler.go -destination ./mocks/baan/baan.handler.go
	mockgen -source ./internal/baan/baan.service.go -destination ./mocks/baan/baan.service.go
	mockgen -source ./internal/selection/selection.handler.go -destination ./mocks/selection/selection.handler.go
	mockgen -source ./internal/selection/selection.service.go -destination ./mocks/selection/selection.service.go
	mockgen -source ./internal/router/context.go -destination ./mocks/router/context.mock.go
	mockgen -source ./internal/validator/validator.go -destination ./mocks/validator/validator.mock.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/isd-sgcu/rpkm67-go-proto@latest

swagger:
	swag init -d ./internal/file -g ../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal