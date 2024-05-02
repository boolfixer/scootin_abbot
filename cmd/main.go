package main

import (
	"github.com/joho/godotenv"
	"main/internal/container"
	"main/internal/http_server"
	"main/internal/simulator"
	"os"
)

func main() {
	loadEnvVars()

	c := container.Bootstrap()
	err := c.Invoke(run)

	if err != nil {
		panic(err)
	}
}

func run(s *http_server.HttpServer) {
	simulateClients()
	s.Serve()
}

func simulateClients() {
	clients := []simulator.ClientSimulator{
		simulator.NewClientSimulator(os.Getenv("USER_API_KEY_1")),
		simulator.NewClientSimulator(os.Getenv("USER_API_KEY_2")),
		simulator.NewClientSimulator(os.Getenv("USER_API_KEY_3")),
	}

	for _, client := range clients {
		go client.Run()
	}
}

func loadEnvVars() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file: %s")
	}
}
