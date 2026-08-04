package repository

import (
	"gorm.io/gorm"
)

type Student struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	ClassID  uint   `json:"class_id"`
	ParentID uint   `json:"parent_id"`
}

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) GetStudentsByClass(classID uint) ([]Student, error) {
	var students []Student
	err := r.db.Where("class_id = ?", classID).
		Preload("Class").
		Preload("Parent").
		Find(&students).Error
	return students, err
}