package main

import (
	"log"
	"os"

	///
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	///
	todotrpt "gotodolist20221123/module/item/transport"
	membertrpt "gotodolist20221123/module/member/transport"
	wallettrpt "gotodolist20221123/module/wallet/transport"
)

func main() {
	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		// log.Fatalln("Missing MySQL connection string.")
		log.Println("Missing MySQL connection string.")
		mysqlConnStr = "root:root@tcp(127.0.0.1:8001)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
	}

	dsn := mysqlConnStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// todo_items
		v1.POST("/items", todotrpt.HanleCreateItem(db))          // create item
		v1.GET("/items", todotrpt.HandleListItems(db))           // list items
		v1.GET("/items/:id", todotrpt.HandleFindItem(db))        // get an item by ID
		v1.PUT("/items/:id", todotrpt.HandleUpdateItem(db))      // edit an item by ID
		v1.DELETE("/items/:id", todotrpt.HandleDeleteAnItem(db)) // delete an item by ID
		// wallets
		v1.POST("/wallets", wallettrpt.HanleCreateItem(db))        // create item
		v1.GET("/wallets", wallettrpt.HandleListItems(db))         // list items
		v1.GET("/wallets/:id", wallettrpt.HandleFindItem(db))      // get an item by ID
		v1.PUT("/wallets/:id", wallettrpt.HandleUpdateItem(db))    // edit an item by ID
		v1.DELETE("/wallets/:id", wallettrpt.HandleDeleteItem(db)) // delete an item by ID
		// members
		v1.POST("/members", membertrpt.HanleCreateItem(db))              // create item
		v1.GET("/members", membertrpt.HandleListItems(db))               // list items
		v1.GET("/members/:username", membertrpt.HandleFindItem(db))      // get an item by username
		v1.PUT("/members/:username", membertrpt.HandleUpdateItem(db))    // edit an item by username
		v1.DELETE("/members/:username", membertrpt.HandleDeleteItem(db)) // delete an item by username
	}

	router.Run()
}

// func createItem(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataItem ToDoItem

// 		// nhận các data từ request
// 		if err := c.ShouldBind(&dataItem); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// preprocess title - trim all spaces
// 		dataItem.Title = strings.TrimSpace(dataItem.Title)

// 		// kiểm tra tính hợp lệ
// 		if dataItem.Title == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
// 			return
// 		}

// 		// do not allow "finished" status when creating a new task
// 		dataItem.Status = "Doing" // set to default

// 		// thao tác với DB
// 		if err := db.Create(&dataItem).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// trả về client với JSON Format
// 		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
// 	}
// }

// func readItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataItem ToDoItem

// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": dataItem})
// 	}
// }

// func getListOfItems(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		type DataPaging struct {
// 			Page  int   `json:"page" form:"page"`
// 			Limit int   `json:"limit" form:"limit"`
// 			Total int64 `json:"total" form:"-"`
// 		}

// 		var paging DataPaging

// 		if err := c.ShouldBind(&paging); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if paging.Page <= 0 {
// 			paging.Page = 1
// 		}

// 		if paging.Limit <= 0 {
// 			paging.Limit = 10
// 		}

// 		offset := (paging.Page - 1) * paging.Limit

// 		var result []ToDoItem

// 		if err := db.Table(ToDoItem{}.TableName()).
// 			Count(&paging.Total).
// 			Offset(offset).
// 			Order("id desc").
// 			Find(&result).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": result})
// 	}
// }

// func editItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var dataItem ToDoItem

// 		if err := c.ShouldBind(&dataItem); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }

// func deleteItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Table(ToDoItem{}.TableName()).
// 			Where("id = ?", id).
// 			Delete(nil).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }
