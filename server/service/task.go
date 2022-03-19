package service

import (
	"fmt"

	"github.com/be3/go_vue_todo/server/model"
)

type TaskService struct{}

func (TaskService) GetTaskList(limit int) (tasks []model.Task) {
	rows, err := Db.Query("select id, content from tasks limit $1", limit)

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
	fmt.Println("SetTask")

	stmt, err := Db.Prepare("insert into tasks (content) values (?)")
	if err != nil {
		fmt.Println("Prepare error")
	}
	result, err := stmt.Exec(task.Content)
	if err != nil {
		fmt.Println("Exec error")
	}
	defer stmt.Close()

	id, _ := result.LastInsertId()
	task.Id = int(id)

	fmt.Println(task)
	return
}
