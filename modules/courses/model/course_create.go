package coursemodel

import "time"

type CourseCreate struct {
	ID        int        `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Name      string     `json:"name" gorm:"column:name;"`
}

func (CourseCreate) TableName() string {
	return "courses"
}
