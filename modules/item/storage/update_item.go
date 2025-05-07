package storage

import (
	"golang.org/x/net/context"
	"social_todo_list.com/modules/item/model"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
