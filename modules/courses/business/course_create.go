package bizcourse

import (
	coursemodel "Gin/modules/courses/model"
	"context"
	"errors"
)

type CreateStore interface {
	Create(ctx context.Context, data *coursemodel.CourseCreate) error
}

type createCourseBiz struct {
	store CreateStore
}

func NewCreateCourseBiz(store CreateStore) *createCourseBiz {
	return &createCourseBiz{store: store}
}

func (biz *createCourseBiz) CreateCourse(ctx context.Context, data *coursemodel.CourseCreate) error {
	if data.Name == "" {
		return errors.New("ten dang nhap khong hop le")
	}

	err := biz.store.Create(ctx, data)

	return err
}
