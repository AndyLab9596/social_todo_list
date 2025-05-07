package biz

import (
	"context"

	"social_todo_list.com/modules/item/model"
)

type DeleteItemByIdStorage interface {
	FindItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemByIdBiz struct {
	store DeleteItemByIdStorage
}

func NewDeleteItemByIdBiz(store DeleteItemByIdStorage) *deleteItemByIdBiz {
	return &deleteItemByIdBiz{store: store}
}

func (biz *deleteItemByIdBiz) DeletetemById(ctx context.Context, id int) error {
	data, err := biz.store.FindItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
