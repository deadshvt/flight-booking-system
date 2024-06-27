include .env
export

help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

proto-gateway: ## Generate gRPC and Go code from gateway proto files
	@protoc -I=. \
			--go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			--grpc-gateway_out=. --grpc-gateway_opt=generate_unbound_methods=true,paths=source_relative \
			gateway/proto/gateway.proto
.PHONY: proto-gateway

proto-ticket: ## Generate gRPC and Go code from ticket proto files
	@protoc -I . \
    		--go_out=. --go_opt=paths=source_relative \
    		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			ticket-service/proto/ticket.proto
.PHONY: proto-ticket

proto-bonus: ## Generate gRPC and Go code from bonus proto files
	@protoc -I . \
    		--go_out=. --go_opt=paths=source_relative \
    		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			bonus-service/proto/bonus.proto
.PHONY: proto-bonus

proto-flight: ## Generate gRPC and Go code from flight proto files
	@protoc -I . \
			--go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			flight-service/proto/flight.proto
.PHONY: proto-flight

proto: proto-gateway proto-ticket proto-bonus proto-flight ## Generate gRPC and Go code from proto files
.PHONY: proto

clean:
	rm -rf gateway/proto/*.go ticket-service/proto/*.go bonus-service/proto/*.go flight-service/proto/*.go
.PHONY: clean

run-gateway: ## Run the gateway
	go run gateway/cmd/gateway/main.go
.PHONY: run-gateway

run-ticket: ## Run the ticket service
	go run ticket-service/cmd/ticket-service/main.go
.PHONY: run-ticket

run-bonus: ## Run the bonus service
	go run bonus-service/cmd/bonus-service/main.go
.PHONY: run-bonus

run-flight: ## Run the flight service
	go run flight-service/cmd/flight-service/main.go
.PHONY: run-flight
