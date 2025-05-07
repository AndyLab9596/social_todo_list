package storage

import (
	"golang.org/x/net/context"
	"social_todo_list.com/modules/item/model"
)

func (s *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.TodoItem{}.
		TableName()).Where(cond).
		Updates(map[string]interface{}{"status": model.ItemStatusDeleted}).Error; err != nil {
		return err
	}

	return nil
}
