package main

import (
	"main/internal/container"
	"main/internal/http_server"
)

func main() {
	c := container.Bootstrap()

	err := c.Invoke(run)

	if err != nil {
		panic(err)
	}
}

func run(s *http_server.HttpServer) {
	s.Serve()
}
