package file

import (
	"io"

	d "github.com/katin-dev/gallery/app/domain/file"
)

type FileRestService struct {
	repo d.FileRepository
}

type FileList struct {
	files []FileDto
	total int
}

func (s *FileRestService) UploadFile(f io.Reader, name string) FileDto {
	// я куда-то загрузил контент...
	file := d.File{
		Name: name,
	}

	s.repo.Create(&file)

	return NewFileDto(file)
}

func (s *FileRestService) getList() FileList {
	files := s.repo.FindBy()

	dtos := make([]FileDto, len(files))
	for i, file := range files {
		dtos[i] = NewFileDto(file)
	}

	return FileList{dtos, 10}
}
