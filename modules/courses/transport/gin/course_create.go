package gincourse

import (
	bizcourse "Gin/modules/courses/business"
	coursemodel "Gin/modules/courses/model"
	coursestorage "Gin/modules/courses/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var newCourse coursemodel.CourseCreate

		if err := ctx.ShouldBind(&newCourse); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := coursestorage.NewSQLStore(db)

		biz := bizcourse.NewCreateCourseBiz(store)

		if err := biz.CreateCourse(ctx.Request.Context(), &newCourse); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": newCourse,
		})

	}
}
