package gincourse

import (
	bizcourse "Gin/modules/courses/business"
	coursestorage "Gin/modules/courses/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListCourse(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		store := coursestorage.NewSQLStore(db)
		biz := bizcourse.NewListCourseBiz(store)

		result, err := biz.ListCourse(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})

	}
}
