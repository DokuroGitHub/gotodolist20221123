package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/member/business"
	"gotodolist20221123/module/member/model"
	"gotodolist20221123/module/member/storage"
)

func HandleListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		limit, limitErr := strconv.Atoi(params.Get("limit"))
		page, pageErr := strconv.Atoi(params.Get("page"))
		where := params.Get("where")
		order := params.Get("order")

		var paging model.DataPaging

		_condition := map[string]interface{}{}
		_order := map[string]bool{}

		if limitErr != nil && limit > -1 {
			paging.Limit = limit
		}
		if pageErr != nil && page > 0 {
			paging.Page = page
		}
		if where != "" {
			_temp := map[string]interface{}{}
			if err := json.Unmarshal([]byte(where), &_temp); err == nil {
				_condition = _temp
			}
		}
		if order != "" {
			_temp := map[string]bool{}
			if err := json.Unmarshal([]byte(order), &_temp); err == nil {
				_order = _temp
			}
		}

		fmt.Println(order)
		fmt.Println(_order)
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := storage.NewMySQLStorage(db)
		business := business.NewListItemsBusiness(storage)

		result, limitErr := business.ListItems(c.Request.Context(), _condition, &paging, _order)

		if limitErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": limitErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
