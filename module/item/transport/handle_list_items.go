package todotrpt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	todoitembusiness "gotodolist20221123/module/item/business"
	todomodel "gotodolist20221123/module/item/model"
	todostorage "gotodolist20221123/module/item/storage"
)

func HandleListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging todomodel.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := todostorage.NewMySQLStorage(db)
		biz := todoitembusiness.NewListToDoItemBiz(storage)

		result, err := biz.ListItems(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
