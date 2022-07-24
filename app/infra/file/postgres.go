package file

import (
	"fmt"

	d "github.com/katin-dev/gallery/app/domain/file"
	"gorm.io/gorm"
)

type PostgresFileRepository struct {
	db *gorm.DB
}

func NewPostgresFileRepository(db *gorm.DB) *PostgresFileRepository {
	return &PostgresFileRepository{db}
}

func (r *PostgresFileRepository) Create(f *d.File) error {
	res := r.db.Create(f)
	if res.Error != nil {
		return fmt.Errorf("Fail to create File: %e", res.Error)
	}

	return nil
}

func (r *PostgresFileRepository) FindBy() ([]d.File, error) {
	var files []d.File
	res := r.db.Order("id DESC").Find(&files)
	if res.Error != nil {
		return nil, fmt.Errorf("Failed to find files: %e", res.Error)
	}

	return files, nil
}
