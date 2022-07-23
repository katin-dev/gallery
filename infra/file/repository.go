package file

import d "github.com/katin-dev/gallery/domain/file"

type FileRepository struct {
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) FindBy() []d.File {
	var files []d.File

	for i := 0; i < 10; i++ {
		files = append(files, d.File{i + 1, "One"})
	}

	return files
}
