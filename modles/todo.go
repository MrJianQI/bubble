package modles

import (
	"bubble/dao"
	"github.com/jinzhu/gorm"
)

//TODO模型
type Todo struct {
	gorm.Model
	Identity int    `json:"id" gorm:"column:identity;type:int(10);"`
	Title    string `json:"title" gorm:"column:title;type:varchar(100);"`
	Status   bool   `json:"status" gorm:"column:status;type:tinyint(1);"`
}

func TableName() string {
	return "todo"
}

//Todo 增删改查
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error //只有一个返回值，不用判断error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("identity=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("identity=?", id).Delete(&Todo{}).Error
	return
}
