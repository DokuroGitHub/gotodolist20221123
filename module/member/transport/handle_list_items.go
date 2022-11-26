package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
		sort := params.Get("sort")

		var _paging model.DataPaging
		var _condition map[string]interface{}
		var _order string

		if limitErr != nil && limit > -1 {
			_paging.Limit = limit
		}
		if pageErr != nil && page > 0 {
			_paging.Page = page
		}
		if where != "" {
			_temp := map[string]interface{}{}
			if err := json.Unmarshal([]byte(where), &_temp); err == nil {
				_condition = _temp
			}
		}
		if sort != "" {
			_order = strings.ReplaceAll(sort, ":", " ")
		}

		if err := c.ShouldBind(&_paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_paging.Process()

		storage := storage.NewMySQLStorage(db)
		business := business.NewListItemsBusiness(storage)

		result, limitErr := business.ListItems(c.Request.Context(), _condition, &_paging, _order)

		if limitErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": limitErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": _paging})
	}
}
