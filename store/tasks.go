package store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/richardimaoka/go-sandbox/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db *sqlx.DB,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT 
        id, user_id, title,
        status, created, modified 
      FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}
