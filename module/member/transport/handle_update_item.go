package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/member/business"
	"gotodolist20221123/module/member/model"
	"gotodolist20221123/module/member/storage"
)

func HandleUpdateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username empty"})
			return
		}

		var dataItem model.Member

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := storage.NewMySQLStorage(db)
		business := business.NewUpdateItemBusiness(storage)

		if err := business.UpdateItem(c.Request.Context(), map[string]interface{}{"username": username}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
