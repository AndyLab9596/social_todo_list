package model

import (
	"errors"

	"social_todo_list.com/common"
)

var (
	ErrTitleIsBlank  = errors.New("title cannot be blank")
	ErrItemIsDeleted = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	// Id          int         `json:"id" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
	// CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at"`
	// UpdatedAt   *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int        `json:"-" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

// *string để update được chuỗi rỗng
type TodoItemUpdate struct {
	Title       string  `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      string  `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
