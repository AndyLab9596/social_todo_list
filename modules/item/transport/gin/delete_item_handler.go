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

func DeleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)
		business := biz.NewDeleteItemByIdBiz(storage)

		if err := business.DeletetemById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
