package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/member/business"
	"gotodolist20221123/module/member/storage"
)

func HandleFindItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username empty"})
			return
		}
		storage := storage.NewMySQLStorage(db)
		business := business.NewFindItemBusiness(storage)

		data, err := business.FindItem(c.Request.Context(), map[string]interface{}{"username": username})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
