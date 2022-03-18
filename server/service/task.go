package service

import "github.com/be3/go_vue_todo/server/model"

type TaskService struct{}

func (TaskService) GetTaskList(limit int) (tasks []model.Task) {
	rows, err := Db.Query("select id, content from tasks limit $1", limit)

	if err != nil {
		return nil
	}

	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Content)
		if err != nil {
			return nil
		}
		tasks = append(tasks, task)
	}
	rows.Close()

	return
}

func (TaskService) SetTask(task *model.Task) (err error) {
	statement := "insert into tasks (content) values ($1) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(task.Content).Scan(&task.Id)
	return
}
