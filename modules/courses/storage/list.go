package coursestorage

import (
	coursemodel "Gin/modules/courses/model"
	"context"
)

func (s *sqlStore) List(context context.Context) ([]coursemodel.CourseCreate, error) {
	var courses []coursemodel.CourseCreate

	if err := s.db.Find(&courses).Error;err != nil {
		return nil,err
	}

	return courses,nil

}
