package ginItem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"social_todo_list.com/common"
	"social_todo_list.com/modules/item/biz"
	"social_todo_list.com/modules/item/storage"
)

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		buisness := biz.NewGetItemByIdBiz(store)

		data, err := buisness.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
