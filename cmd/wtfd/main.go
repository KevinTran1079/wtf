package main

import (
	"context"
	"fmt"
	"log"

	"github.com/KevinTran1079/wtf/httpserver"
	"github.com/KevinTran1079/wtf/inmem"
)

type Main struct {
	HTTPServer *httpserver.Server
}

func main() {
	m := NewMain()
	if err := m.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func NewMain() *Main {
	taskService := inmem.NewTaskService()

	return &Main{
		HTTPServer: httpserver.NewServer(taskService),
	}
}

func (m *Main) Run(ctx context.Context) error {
	fmt.Println("Starting server and listening on port 8080")
	return m.HTTPServer.ListenAndServe()
}
