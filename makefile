.DEFAULT: server

server:
	go run -tags=dynamic cmd/server.go

run:
	go run -tags=dynamic cmd/main.go && ffplay out.png
