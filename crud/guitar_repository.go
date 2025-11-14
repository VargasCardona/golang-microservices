package crud

import (
	"context"

	"gorm.io/gorm"
)

type GuitarRepository interface {
	Create(ctx context.Context, g *GuitarModel) error
	GetByID(ctx context.Context, id uint) (*GuitarModel, error)
	List(ctx context.Context) ([]GuitarModel, error)
	Update(ctx context.Context, g *GuitarModel) error
	Delete(ctx context.Context, id uint) error
}

type guitarRepository struct {
	db *gorm.DB
}

func NewGuitarRepository(db *gorm.DB) GuitarRepository {
	return &guitarRepository{db: db}
}

// CREATE
func (r *guitarRepository) Create(ctx context.Context, g *GuitarModel) error {
	return r.db.WithContext(ctx).Create(g).Error
}

// READ ONE
func (r *guitarRepository) GetByID(ctx context.Context, id uint) (*GuitarModel, error) {
	var g GuitarModel
	if err := r.db.WithContext(ctx).First(&g, id).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

// READ LIST
func (r *guitarRepository) List(ctx context.Context) ([]GuitarModel, error) {
	var guitars []GuitarModel
	if err := r.db.WithContext(ctx).Find(&guitars).Error; err != nil {
		return nil, err
	}
	return guitars, nil
}

// UPDATE
func (r *guitarRepository) Update(ctx context.Context, g *GuitarModel) error {
	return r.db.WithContext(ctx).Save(g).Error
}

// DELETE
func (r *guitarRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&GuitarModel{}).Error
}
