package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gotodolist20221123/module/wallet/business"
	"gotodolist20221123/module/wallet/model"
	"gotodolist20221123/module/wallet/storage"
)

func HandleListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		fmt.Println(params)
		limit, limitErr := strconv.Atoi(params.Get("limit"))
		page, pageErr := strconv.Atoi(params.Get("page"))
		where := params.Get("where")
		not := params.Get("not")
		or := params.Get("or")
		sort := params.Get("sort")

		_paging := model.DataPaging{}
		_where := map[string]interface{}{}
		_not := map[string]interface{}{}
		_or := map[string]interface{}{}
		_order := ""

		if limitErr != nil && limit > -1 {
			_paging.Limit = limit
		}
		if pageErr != nil && page > 0 {
			_paging.Page = page
		}
		if where != "" {
			// in, nin, neq, gt, gte, lt, lte
			// where={"id":"gt:1"}
			// where=["id=in:(1,2,3,4)"]
			_temp := map[string]interface{}{}
			if err := json.Unmarshal([]byte(where), &_temp); err == nil {
				for k, v := range _temp {
					_where[k] = v
				}
			}
			// if err := json.Unmarshal([]byte(where), &_temp); err == nil {
			// 	for _, v := range _temp {
			// 		v = strings.ReplaceAll(v, "=in:", " IN ")
			// 		v = strings.ReplaceAll(v, "=nin:", " NOT IN ")
			// 		v = strings.ReplaceAll(v, "=neq:", " <> ")
			// 		v = strings.ReplaceAll(v, "=gt:", " > ")
			// 		v = strings.ReplaceAll(v, "=gte:", " >= ")
			// 		v = strings.ReplaceAll(v, "=lt:", " < ")
			// 		v = strings.ReplaceAll(v, "=lte:", " <= ")
			// 		_where = append(_where, v)
			// 	}
			// }
		}
		fmt.Println(where)
		fmt.Println(_where)
		if not != "" {
			_temp := map[string]interface{}{}
			if err := json.Unmarshal([]byte(not), &_temp); err == nil {
				for k, v := range _temp {
					_not[k] = v
				}
			}
		}
		if or != "" {
			_temp := map[string]interface{}{}
			if err := json.Unmarshal([]byte(or), &_temp); err == nil {
				for k, v := range _temp {
					_or[k] = v
				}
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

		result, limitErr := business.ListItems(c.Request.Context(), _where, _not, _or, &_paging, _order)

		if limitErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": limitErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": _paging})
	}
}
