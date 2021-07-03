package bizcourse

import (
	coursemodel "Gin/modules/courses/model"
	"context"
)

type ListStore interface {
	List(ctx context.Context) ([]coursemodel.CourseCreate, error)
}

type listCourseBiz struct {
	store ListStore
}

func NewListCourseBiz(store ListStore) *listCourseBiz {
	return &listCourseBiz{
		store: store,
	}
}

func (biz *listCourseBiz) ListCourse(ctx context.Context) ([]coursemodel.CourseCreate, error) {
	result, err := biz.store.List(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
