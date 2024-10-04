package repositories

import (
	"gorm.io/gorm"
)

type Model interface {
	any
}

type BaseRepository[T Model] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new BaseRepository with the given gorm.DB
func NewBaseRepository[T Model](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// Common methods for all repositories
func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var result T
	err := r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var result []T
	err := r.db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *BaseRepository[T]) Create(model *T) error {
	return r.db.Create(model).Error
}

func (r *BaseRepository[T]) Update(model *T) error {
	return r.db.Save(model).Error
}

func (r *BaseRepository[T]) Delete(model *T) error {
	return r.db.Delete(model).Error
}
