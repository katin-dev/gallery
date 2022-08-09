package file

import (
	"context"
	"fmt"
	"strconv"

	d "github.com/katin-dev/gallery/app/domain/file"
	"github.com/minio/minio-go/v7"
)

type FileRestService struct {
	repo     d.FileRepository
	s3client *minio.Client
	bucket   string
}

type FileList struct {
	files []*FileDto
	total int64
}

func NewFileRestService(fileRepository d.FileRepository, s3client *minio.Client, bucket string) *FileRestService {
	return &FileRestService{fileRepository, s3client, bucket}
}

func (s *FileRestService) UploadFile(filePath string, name string, size int64) (*FileDto, error) {
	// я куда-то загрузил контент...
	file := d.File{
		Name: name,
	}

	s.repo.Create(&file)

	fileKey := strconv.Itoa(int(file.ID))

	_, err := s.s3client.FPutObject(context.Background(), s.bucket, fileKey, filePath, minio.PutObjectOptions{
		ContentType: "application/csv",
	})
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to upload file %s to AWS S3: %e", name, err)
	}

	return NewFileDto(file), nil
}

func (s *FileRestService) getList() (*FileList, error) {
	files, err := s.repo.FindBy()
	if err != nil {
		return nil, err
	}
	total, err := s.repo.CountBy()
	if err != nil {
		return nil, err
	}

	dtos := make([]*FileDto, len(files))
	for i, file := range files {
		dtos[i] = NewFileDto(file)
	}

	return &FileList{dtos, total}, nil
}
