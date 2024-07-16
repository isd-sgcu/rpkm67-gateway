pull-latest-mac:
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/firstdate-rpkm67-frontend:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-gateway:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-auth:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-backend:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-checkin:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-store:latest

pull-latest-windows:
	docker pull ghcr.io/isd-sgcu/rpkm67-gateway:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-auth:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-backend:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-checkin:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-store:latest

docker:
	docker rm -v -f $$(docker ps -qa) || echo "No containers found. Skipping removal."
	docker-compose up

docker-qa:
	docker rm -v -f $$(docker ps -qa) || echo "No containers found. Skipping removal."
	docker-compose -f docker-compose.qa.yml up

server:
	go run cmd/main.go

watch: 
	air

mock-gen:
	mockgen -source ./internal/auth/auth.service.go -destination ./mocks/auth/auth.service.go
	mockgen -source ./internal/pin/pin.service.go -destination ./mocks/pin/pin.service.go
	mockgen -source ./internal/group/group.service.go -destination ./mocks/group/group.service.go
	mockgen -source ./internal/selection/selection.service.go -destination ./mocks/selection/selection.service.go
	mockgen -source ./internal/selection/selection.client.go -destination ./mocks/selection/selection.client.go
	mockgen -source ./internal/checkin/checkin.service.go -destination ./mocks/checkin/checkin.service.go
	mockgen -source ./internal/checkin/checkin.client.go -destination ./mocks/checkin/checkin.client.go
	mockgen -source ./internal/object/object.service.go -destination ./mocks/object/object.service.go
	mockgen -source ./internal/object/object.client.go -destination ./mocks/object/object.client.go
	mockgen -source ./internal/stamp/stamp.service.go -destination ./mocks/stamp/stamp.service.go
	mockgen -source ./internal/stamp/stamp.client.go -destination ./mocks/stamp/stamp.client.go
	mockgen -source ./internal/user/user.service.go -destination ./mocks/user/user.service.go
	mockgen -source ./internal/user/user.client.go -destination ./mocks/user/user.client.go
	mockgen -source ./internal/context/context.go -destination ./mocks/context/context.mock.go
	mockgen -source ./internal/validator/validator.go -destination ./mocks/validator/validator.mock.go
	mockgen -source ./internal/group/group.handler.go -destination ./mocks/group/group.handler.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/isd-sgcu/rpkm67-go-proto@latest

model:
	go get github.com/isd-sgcu/rpkm67-model@latest

swagger:
	swag init -d ./internal -g ../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal
