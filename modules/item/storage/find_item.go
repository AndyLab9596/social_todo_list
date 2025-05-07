package storage

import (
	"context"

	"social_todo_list.com/modules/item/model"
)

func (s *sqlStorage) FindItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
