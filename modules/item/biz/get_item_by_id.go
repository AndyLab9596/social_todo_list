package biz

import (
	"context"

	"social_todo_list.com/modules/item/model"
)

type GetItemStorage interface {
	FindItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemByIdBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := biz.store.FindItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return data, nil
}
