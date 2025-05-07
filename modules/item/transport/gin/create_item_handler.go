package ginItem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"social_todo_list.com/common"
	"social_todo_list.com/modules/item/biz"
	"social_todo_list.com/modules/item/model"
	"social_todo_list.com/modules/item/storage"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation

		// bind -> UnmarshalJSON -> bind data to *ItemStatus
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
