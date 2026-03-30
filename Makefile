run app migrate:
		DNV=dev go run ./cmd/app/main.go -migrate
start server:
		DNV=dev go run ./cmd/app/main.go

# .PHONY: run ducker up run docker build run docker logs