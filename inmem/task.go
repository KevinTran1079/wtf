package inmem

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KevinTran1079/wtf"
)

type TaskService struct {
	taskList []*wtf.Task
}

func NewTaskService() *TaskService {
	now := time.Now().UTC()

	return &TaskService{
		taskList: []*wtf.Task{
			{
				ID:          1,
				Title:       "Learn Go handlers",
				Description: "Build a simple REST endpoint",
				Completed:   false,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          2,
				Title:       "Wire the in-memory service",
				Description: "Connect handlers to a task service",
				Completed:   false,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          3,
				Title:       "Return JSON responses",
				Description: "Encode task data through the response writer",
				Completed:   true,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}
}

func (s *TaskService) FindTasks() ([]*wtf.Task, error) {
	return s.taskList, nil
}

func (s *TaskService) CreateTask(task *wtf.Task) (*wtf.Task, error) {
	s.taskList = append(s.taskList, task)
	return task, nil
}

func (s *TaskService) FindTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tasks, err := s.FindTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}
