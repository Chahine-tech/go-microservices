server:
	go run cmd/main.go

docker:
	docker build -t go-server .
	docker run -p 8080:8080 go-server