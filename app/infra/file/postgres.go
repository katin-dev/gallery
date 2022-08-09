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

func (r *PostgresFileRepository) CountBy() (int64, error) {
	var file d.File
	var total int64
	res := r.db.Model(&file).Count(&total)
	if res.Error != nil {
		return 0, fmt.Errorf("Failed to count files: %e", res.Error)
	}

	return total, nil
}
