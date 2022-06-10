package service

import (
	"fmt"
	"log"
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
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
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

	Stmt, err = Db.Prepare("insert into tasks(content, user_id) value(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := Stmt.Exec(task.Content, task.User.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer Stmt.Close()

	id, _ := result.LastInsertId()
	task.Id = int(id)

	fmt.Println(task)
	return
}

func (TaskService) ChangeTaskById(id string, task *model.Task) (err error) {
	fmt.Println("ChangeTask")

	Stmt, err = Db.Prepare("update tasks set content = ? where id = ?")
	if err != nil {
		fmt.Println("Prepare error")
		return
	}
	_, err = Stmt.Exec(task.Content, id)
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

func (TaskService) DeleteTaskById(id string) (err error) {
	fmt.Println("DeleteTask")

	Stmt, err = Db.Prepare("delete from tasks where id = ?")
	if err != nil {
		fmt.Println("Prepare error")
		return
	}
	defer Stmt.Close()

	_, err = Stmt.Exec(id)
	if err != nil {
		fmt.Println("Exec error")
		return
	}

	return
}
