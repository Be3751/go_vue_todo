package service

import (
	"fmt"
	"strconv"

	"github.com/be3/go_vue_todo/server/model"
)

type TaskService struct{}

func (TaskService) GetTaskList() ([]model.Task, error) {
	fmt.Println("GetTaskList")

	var tasks []model.Task
	rows, err := Db.Query("select id, content from tasks")
	if err != nil {
		fmt.Println("Select error")
	}
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		fmt.Println(task)
	}
	rows.Close()

	return tasks, nil
}

func (TaskService) GetTaskById(id string) (model.Task, error) {
	fmt.Println("GetTaskById")

	var task model.Task
	err := Db.QueryRow("select id, content from tasks where id = ?", id).Scan(&task.Id, &task.Content)
	if err != nil {
		fmt.Println("Select error")
	}
	return task, nil
}

func (TaskService) AddTask(task *model.Task) (err error) {
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

func (TaskService) ChangeTaskById(id string, task *model.Task) (err error) {
	fmt.Println("ChangeTask")

	stmt, err := Db.Prepare("update tasks set content = ? where id = ?")
	if err != nil {
		fmt.Println("Prepare error")
		return
	}
	_, err = stmt.Exec(task.Content, id)
	if err != nil {
		fmt.Println("Exec error")
		return
	}
	task.Id, err = strconv.Atoi(id)
	if err != nil {
		fmt.Println("String to Int convert error")
		return
	}

	return
}
