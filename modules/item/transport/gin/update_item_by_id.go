package ginItem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"social_todo_list.com/common"
	"social_todo_list.com/modules/item/biz"
	"social_todo_list.com/modules/item/model"
	"social_todo_list.com/modules/item/storage"
)

func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)
		business := biz.NewUpdateItemByIdBiz(storage)

		business.UpdateItemById(c.Request.Context(), id, &data)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
