package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/item/business"
	"gotodolist20221123/module/item/storage"
)

func HandleDeleteAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// setup dependencies
		storage := storage.NewMySQLStorage(db)
		business := business.NewDeleteItemBusiness(storage)

		if err := business.DeleteItem(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
