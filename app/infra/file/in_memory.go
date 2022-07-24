package file

import d "github.com/katin-dev/gallery/app/domain/file"

type InMemoryFileRepository struct {
	files       []d.File
	fileCounter int
}

func NewFileRepository() *InMemoryFileRepository {
	return &InMemoryFileRepository{
		make([]d.File, 0),
		0,
	}
}

func (r *InMemoryFileRepository) Create(f *d.File) {
	r.fileCounter++
	f.ID = uint(r.fileCounter)
	r.files = append(r.files, *f)
}

func (r *InMemoryFileRepository) FindBy() []d.File {
	return r.files
}
