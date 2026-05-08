package httpserver

import (
	"net/http"

	"github.com/KevinTran1079/wtf/inmem"
)

type Server struct {
	Addr   string
	server *http.Server
	mux    *http.ServeMux

	TaskService *inmem.TaskService
}

func NewServer(taskService *inmem.TaskService) *Server {
	newMux := http.NewServeMux()
	s := &Server{
		Addr: ":8080",
		server: &http.Server{
			Addr:    ":8080",
			Handler: newMux,
		},
		mux:         newMux,
		TaskService: taskService,
	}

	s.RegisterRoutes()
	return s
}

func (s *Server) RegisterRoutes() {
	s.mux.HandleFunc("/tasks", s.TaskService.FindTasksHandler)
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}
