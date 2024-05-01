package main

import (
	"main/internal/container"
	"main/internal/http_server"
	"main/internal/simulator"
)

func main() {
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
		simulator.NewClientSimulator("yLxCMVd*p9hZNnvYfGx$yQezPBE3@8Lp9sgKhE9QwtQMM!v6y%4$Q&UuqJjC"),
		simulator.NewClientSimulator("rcYPJRUGf3xv&w4ny%Qhs1^&LJN@&T@H8$3srvheyYP&XJXs^S@sU1QFQ$Hv"),
		simulator.NewClientSimulator("3JFj84#XE4^18jfbhT2vfB#u3Ev4DrsQ*xrXJg7N5dWEgTh2XmTEx%EQ4D8U"),
	}

	for _, client := range clients {
		go client.Run()
	}
}
