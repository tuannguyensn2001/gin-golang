package coursestorage

import (
	coursemodel "Gin/modules/courses/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context,data *coursemodel.CourseCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}