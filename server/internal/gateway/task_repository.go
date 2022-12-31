package gateway

import (
	"database/sql"
	"fmt"

	"github.com/be3/go_vue_todo/server/internal/domain/model"
	"github.com/be3/go_vue_todo/server/internal/utils"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(d *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: d,
	}
}

func (r *TaskRepository) FindById(taskID string) (*model.Task, error) {
	var task *model.Task
	err := r.db.QueryRow("select id, content, deadline from tasks where id = ?", taskID).Scan(task.Id, task.Content, task.Deadline)
	if err != nil {
		return nil, fmt.Errorf("query row error: %s", err.Error())
	}
	return task, nil
}

func (r *TaskRepository) FindByUserId(userID string) ([]*model.Task, error) {
	stmt := "select id, content, deadline from tasks where user_id = " + userID
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("query error: %s", err.Error())
	}

	var tasks []*model.Task
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Content, &task.Deadline)
		if err != nil {
			return nil, fmt.Errorf("scan error: %s", err.Error())
		}
		tasks = append(tasks, &task)
	}
	rows.Close()

	return tasks, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
	stmt, err := r.db.Prepare("insert into tasks(content, user_id, deadline, created_at) value(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("prepare error: %s", err.Error())
	}

	dateCreated := utils.ExtractDate(task.CreatedAt)
	result, err := Stmt.Exec(task.Content, task.User.Id, task.Deadline, dateCreated)
	if err != nil {
		return fmt.Errorf("exec error: %s", err.Error())
	}
	defer stmt.Close()

	taskID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("get inserted task's ID error: %s", err.Error())
	}
	task.Id = int(taskID)

	return nil
}

func (r *TaskRepository) Update(task *model.Task) error {
	stmt, err := r.db.Prepare("update tasks set content = ?, updated_at = ? where id = ?")
	if err != nil {
		return fmt.Errorf("prepare error: %s", err.Error())
	}

	dateUpdated := utils.ExtractDate(task.UpdatedAt)
	_, err = stmt.Exec(task.Content, dateUpdated, task.Id)
	if err != nil {
		return fmt.Errorf("exec error: %s", err.Error())
	}

	return nil
}

func (r *TaskRepository) Delete(task *model.Task) error {
	stmt, err := r.db.Prepare("delete from tasks where id = ?")
	if err != nil {
		return fmt.Errorf("prepare error: %s", err.Error())
	}
	defer stmt.Close()

	_, err = Stmt.Exec(task.Id)
	if err != nil {
		return fmt.Errorf("exec error: %s", err.Error())
	}

	return nil
}
