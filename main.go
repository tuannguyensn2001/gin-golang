package main

import (
	gincourse "Gin/modules/courses/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type RestaurantCreate struct {
//	SQLModel
//	Name    string `json:"name" gorm:"column:name;"`
//	Address string `json:"address" gorm:"column:addr;"`
//}

func main() {

	dsn := "root:password@tcp(172.17.0.1:3306)/golang?charset=utf8&parseTime=True"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r := gin.Default()

	r.POST("/course", gincourse.CreateCourse(db))

	r.GET("/courses",gincourse.ListCourse(db))

	err := r.Run(":5000")
	if err != nil {
		return
	}
}

//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//type SQLModel struct {
//	ID        int       `json:"id" gorm:"column:id;"`
//	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
//	Status    int       `json:"status" gorm:"column:status;default:1;"`
//}
//
//type Note struct {
//	SQLModel
//	Name       string `gorm:"column:title;"`
//	CategoryId int    `gorm:"column:category_id;"`
//}
//
//func (Note) TableName() string { return "notes" }
//
//type NoteUpdate struct {
//	Name       *string `gorm:"column:title;"`
//	CategoryId *int    `gorm:"column:category_id;"`
//	Status     *int    `gorm:"column:status;"`
//}
//
//func (NoteUpdate) TableName() string { return Note{}.TableName() }
//
//type Restaurant struct {
//	SQLModel
//	Name    string `json:"name" gorm:"column:name;"`
//	Address string `json:"address" gorm:"column:addr;"`
//}
//
//func (Restaurant) TableName() string { return "restaurants" }
//
//type RestaurantCreate struct {
//	SQLModel
//	Name    string `json:"name" gorm:"column:name;"`
//	Address string `json:"address" gorm:"column:addr;"`
//}
//
//func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
//
//type RestaurantUpdate struct {
//	Name    *string `json:"name" gorm:"column:name;"`
//	Address *string `json:"address" gorm:"column:addr;"`
//	Status  *int    `json:"-" gorm:"column:status;"`
//}
//
//func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
//
//func main() {
//	log.Println("Hello world")
//
//		dsn := "root:password@tcp(172.17.0.1:3306)/golang?charset=utf8&parseTime=True"
//		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatalln("Error mysql:", err)
//	}
//
//	db = db.Debug()
//
//	//n := Note{
//	//	Name:       "Note 12",
//	//	CategoryId: 1,
//	//}
//
//	//log.Println(n.ID)
//	//
//	//if err := db.Create(&n).Error; err != nil {
//	//	log.Println(err)
//	//}
//	//
//	//log.Println(n.ID)
//
//	//var myNote Note
//	//
//
//	//}
//	//
//	//log.Println(listNote)
//	//
//	//myNote.Name = "New note 15"
//	//
//	//emptyString := ""
//	//zeroCategoryId := 0
//	//
//	//if err := db.Where("id = ?", 15).
//	//	Updates(NoteUpdate{
//	//		Name:       &emptyString,
//	//		CategoryId: &zeroCategoryId,
//	//	}).Error; err != nil {
//	//	log.Println(err)
//	//}
//
//	//if err := db.Table(myNote.TableName()).Where("id = ?", 15).Delete(nil); err != nil {
//	//	log.Println(err)
//	//}
//
//	r := gin.Default()
//
//	v1 := r.Group("/v1")
//	{
//		restaurants := v1.Group("/restaurants")
//		{
//			// CRUD
//			restaurants.POST("", func(c *gin.Context) {
//				var newRestaurant RestaurantCreate
//
//				if err := c.ShouldBind(&newRestaurant); err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{"data": newRestaurant})
//
//				//if err := db.Create(&newRestaurant).Error; err != nil {
//				//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//				//	return
//				//}
//				//
//				//c.JSON(http.StatusOK, gin.H{"data": newRestaurant.ID})
//			})
//
//			restaurants.GET("", func(c *gin.Context) {
//				var result []Restaurant
//
//				var paging struct {
//					Page  int   `json:"page" form:"page"`
//					Limit int   `json:"limit" form:"limit"`
//					Total int64 `json:"total" form:"total"`
//				}
//
//				if err := c.ShouldBind(&paging); err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				if paging.Limit <= 0 {
//					paging.Limit = 10
//				}
//
//				if paging.Page <= 0 {
//					paging.Page = 1
//				}
//
//				if err := db.Table(Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
//					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//					return
//				}
//
//				if err := db.
//					Limit(paging.Limit).
//					Offset((paging.Page - 1) * paging.Limit).
//					Order("id desc").
//					Find(&result).Error; err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
//			})
//
//			restaurants.PUT("/:id", func(c *gin.Context) {
//				id, err := strconv.Atoi(c.Param("id"))
//
//				if err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				var dataUpdate RestaurantUpdate
//
//				if err := c.ShouldBind(&dataUpdate); err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				if err := db.Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{"data": 1})
//			})
//
//			restaurants.DELETE("/:id", func(c *gin.Context) {
//				id, err := strconv.Atoi(c.Param("id"))
//
//				if err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{"data": 1})
//			})
//
//			restaurants.GET("/:id", func(c *gin.Context) {
//				id, err := strconv.Atoi(c.Param("id"))
//
//				if err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				var data Restaurant
//
//				if err := db.Where("id = ?", id).First(&data).Error; err != nil {
//					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//					return
//				}
//
//				c.JSON(http.StatusOK, gin.H{"data": data})
//			})
//		}
//	}
//
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//
//	r.Run(":5000")
//
//}//if err := db.
//	//	Where("id = ?", 15).
//	//	First(&myNote).Error; err != nil {
//	//	log.Println(err)
//	//}
//	//
//	//log.Println(myNote.Name, myNote.CategoryId)
//	//
//	//var listNote []Note // slice <> array
//	//
//	//if err := db.Find(&listNote); err != nil {
//	//	log.Println(err)
