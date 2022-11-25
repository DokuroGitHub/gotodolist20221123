package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/wallet/business"
	"gotodolist20221123/module/wallet/model"
	"gotodolist20221123/module/wallet/storage"
)

func HandleListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging model.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := storage.NewMySQLStorage(db)
		business := business.NewListItemsBusiness(storage)

		result, err := business.ListItems(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
