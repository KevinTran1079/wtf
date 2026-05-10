package sqlite

import (
	"context"

	"github.com/KevinTran1079/wtf"
)

type TaskService struct {
	db *DB
}

func NewTaskService(db *DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) FindTasks(ctx context.Context) ([]*wtf.Task, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	defer tx.Rollback()

	return findTasks(ctx, tx)
}

func findTasks(ctx context.Context, tx *Tx) (_ []*wtf.Task, n int, err error) {
	rows, err := tx.QueryContext(ctx, `
		SELECT
			id,
			title,
			description,
			completed,
			created_at,
			updated_at,
			COUNT(*) OVER()
		FROM tasks
		ORDER BY id ASC
	`)

	if err != nil {
		return nil, n, err
	}

	defer rows.Close()

	tasks := make([]*wtf.Task, 0)

	for rows.Next() {
		var task wtf.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
			&n,
		); err != nil {
			return nil, 0, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return tasks, n, nil
}
