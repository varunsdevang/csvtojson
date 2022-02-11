run:
	go run main.go

docs:
	rm -rf docs
	swag init -g controller/server.go