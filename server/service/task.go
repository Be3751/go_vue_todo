package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/be3/go_vue_todo/server/model"
	"github.com/be3/go_vue_todo/server/utils"
)

type TaskService struct{}

func (TaskService) GetTaskList(userId string) ([]model.Task, error) {
	fmt.Println("GetTaskList")

	stmt := "select id, content, deadline from tasks where user_id = " + userId
	rows, err := Db.Query(stmt)
	if err != nil {
		fmt.Println("Select error")
		fmt.Println(err)
		return nil, err
	}

	var tasks []model.Task
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Content, &task.Deadline)
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
	err := Db.QueryRow("select id, content, deadline from tasks where id = ?", id).Scan(&task.Id, &task.Content, &task.Deadline)
	if err != nil {
		fmt.Println("Select error")
	}
	return task, nil
}

func (TaskService) AddTask(task *model.Task) (err error) {
	fmt.Println("AddTask")

	Stmt, err = Db.Prepare("insert into tasks(content, user_id, deadline, created_at) value(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	dateCreated := utils.ExtractDate(task.CreatedAt)
	result, err := Stmt.Exec(task.Content, task.User.Id, task.Deadline, dateCreated)
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

	Stmt, err = Db.Prepare("update tasks set content = ?, updated_at = ? where id = ?")
	if err != nil {
		fmt.Println("Prepare error")
		return
	}

	dateUpdated := utils.ExtractDate(task.UpdatedAt)
	_, err = Stmt.Exec(task.Content, dateUpdated, id)
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
