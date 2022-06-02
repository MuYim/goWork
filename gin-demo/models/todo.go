package models

import "gin-demo/dao"

/**
create table `todos`(
	`id` int NOT NULL AUTO_INCREMENT primary key,
	`title` varchar(1000) NOT NULL,
	`status` boolean not null default 0
)engine=InnoDB default charset=utf8mb4;
*/
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo的增删改查放到model中
*/

//增
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

//删
func DeleteTodoById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}

//改
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return err
}

//查
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}
