package storage

import (
	"golang.org/x/net/context"
	"social_todo_list.com/modules/item/model"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	// *ItemStatus -> Value()
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
