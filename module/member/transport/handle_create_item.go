package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/member/business"
	"gotodolist20221123/module/member/model"
	"gotodolist20221123/module/member/storage"
)

func HanleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem model.Member

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// setup dependencies
		storage := storage.NewMySQLStorage(db)
		business := business.NewCreateItemBusiness(storage)

		if err := business.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.UserName})
	}
}
