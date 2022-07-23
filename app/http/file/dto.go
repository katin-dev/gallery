package file

import "github.com/katin-dev/gallery/app/domain/file"

type FileDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewFileDto(f file.File) FileDto {
	dto := FileDto{
		uint(f.Id), f.Name,
	}

	return dto
}
