BINARY_NAME = auth

PROTO_PATH = api/auth
OUT_PATH = pkg/auth

.PHONY: generate-pb
generate-pb:
	@mkdir -p ${OUT_PATH}
	@protoc --proto_path=${PROTO_PATH} \
	--go_out=${OUT_PATH} --go_opt=paths=source_relative \
	--go-grpc_out=${OUT_PATH} --go-grpc_opt=paths=source_relative \
	api/auth/auth.proto
	@echo generation complete

.PHONY: build
build:
	@go build -o bin/${BINARY_NAME} cmd/auth/main.go

.PHONY: run
run: build
	@./bin/${BINARY_NAME}

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: compose-build
compose-build:
	@docker compose build

.PHONY: compose-up
compose-up:
	@docker compose up

.PHONY: compose-down
compose-down:
	@docker compose down