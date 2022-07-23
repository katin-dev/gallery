package file

import d "github.com/katin-dev/gallery/app/domain/file"

type FileRepository struct {
	files       []d.File
	fileCounter int
}

func NewFileRepository() *FileRepository {
	return &FileRepository{
		make([]d.File, 0),
		0,
	}
}

func (r *FileRepository) Create(f *d.File) {
	r.fileCounter++
	f.Id = r.fileCounter
	r.files = append(r.files, *f)
}

func (r *FileRepository) FindBy() []d.File {
	return r.files
}
