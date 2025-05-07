package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"social_todo_list.com/common"
	"social_todo_list.com/modules/item/model"
	ginItem "social_todo_list.com/modules/item/transport/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) / v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// PUT /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delet item detail by id)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")

		{
			items.POST("", ginItem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginItem.GetItem(db))
			items.PATCH("/:id", ginItem.UpdateItem(db))
			items.DELETE("/:id", ginItem.DeleteItem(db))
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.
			Order("id desc").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))

	}
}
