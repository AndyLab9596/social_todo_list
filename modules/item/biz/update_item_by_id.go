package biz

import (
	"context"

	"social_todo_list.com/modules/item/model"
)

type UpdateItemByIdStorage interface {
	FindItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemByIdBiz struct {
	store UpdateItemByIdStorage
}

func NewUpdateItemByIdBiz(store UpdateItemByIdStorage) *updateItemByIdBiz {
	return &updateItemByIdBiz{store: store}
}

func (biz *updateItemByIdBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.store.FindItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIsDeleted
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
