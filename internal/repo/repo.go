package repo

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/masputrawae/todo-cli/internal/model"
)

// Errors message
var (
	ErrIDCantEmpty    = errors.New("id tidak boleh kosong")
	ErrTaskCantEmpty  = errors.New("tugas tidak boleh kosong")
	ErrTaskNotFound   = errors.New("tugas tidak ditemukan")
	ErrNothingUpdated = errors.New("tidak ada yang diperbarui")
	ErrInputCantNil   = errors.New("input tidak boleh nil")
)

// Create input types
type CreateInputType struct {
	Status   *string
	Priority *string
	Category *string
	Task     *string
}

// Update inpit types
type UpdateInputType struct {
	ID       *int
	Status   *string
	Priority *string
	Category *string
	Task     *string
}

// Find input types
type FindInputType struct {
	ID       *int
	Status   *string
	Priority *string
	Category *string
}

// Delete input types
type DeleteInputType struct {
	ID *int
}

type Repo struct {
	DB *sql.DB
}

type Manage interface {
	Create(ctx context.Context, input *CreateInputType) error
	Update(ctx context.Context, input *UpdateInputType) error
	Delete(ctx context.Context, input *DeleteInputType) error
	Find(ctx context.Context, input *FindInputType) ([]model.Todo, error)
}

func New(db *sql.DB) Manage {
	return &Repo{DB: db}
}

func (r *Repo) Create(ctx context.Context, input *CreateInputType) error {
	if input == nil {
		return ErrInputCantNil
	}

	// check input task
	if input.Task == nil {
		return ErrTaskCantEmpty
	}

	var cols []string
	var vals []any

	cols = append(cols, "task")
	vals = append(vals, *input.Task)

	// check input status
	if input.Status != nil {
		cols = append(cols, "status")
		vals = append(vals, *input.Status)
	}

	// check input priority
	if input.Priority != nil {
		cols = append(cols, "priority")
		vals = append(vals, *input.Priority)
	}

	// check input category
	if input.Category != nil {
		cols = append(cols, "category")
		vals = append(vals, *input.Category)
	}

	// build query sql
	query, args, err := sq.Insert("todos").Columns(cols...).Values(vals...).ToSql()
	if err != nil {
		return err
	}

	// insert to database
	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r *Repo) Update(ctx context.Context, input *UpdateInputType) error {
	if input == nil {
		return ErrInputCantNil
	}

	if input.ID == nil {
		return ErrIDCantEmpty
	}
	updates := map[string]any{}

	// check task
	if input.Task != nil {
		updates["task"] = *input.Task
	}

	// check status
	if input.Status != nil {
		updates["status"] = *input.Status
	}

	// check priority
	if input.Priority != nil {
		updates["priority"] = *input.Priority
	}

	// check category
	if input.Category != nil {
		updates["category"] = *input.Category
	}

	// check updates
	if len(updates) == 0 {
		return ErrNothingUpdated
	}

	updates["lastmod"] = sq.Expr("CURRENT_TIMESTAMP")
	query, args, err := sq.Update("todos").SetMap(updates).Where(sq.Eq{"id": *input.ID}).ToSql()
	if err != nil {
		return err
	}

	result, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *Repo) Delete(ctx context.Context, input *DeleteInputType) error {
	if input == nil {
		return ErrInputCantNil
	}
	if input.ID == nil {
		return ErrIDCantEmpty
	}

	query, args, err := sq.Delete("todos").Where(sq.Eq{"id": *input.ID}).ToSql()
	if err != nil {
		return err
	}
	result, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrTaskNotFound
	}
	return nil
}

func (r *Repo) Find(ctx context.Context, input *FindInputType) ([]model.Todo, error) {
	var todos []model.Todo
	filters := map[string]any{}

	if input.ID != nil {
		filters["id"] = *input.ID
	}
	if input.Status != nil {
		filters["status"] = *input.Status
	}
	if input.Priority != nil {
		filters["priority"] = *input.Priority
	}
	if input.Category != nil {
		filters["category"] = *input.Category
	}

	builder := sq.Select("id", "status", "priority", "category", "task", "lastmod").From("todos")
	if len(filters) > 0 {
		builder = builder.Where(sq.Eq(filters))
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Status, &t.Priority, &t.Category, &t.Task, &t.Lastmod); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}
