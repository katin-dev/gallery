package file

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	d "github.com/katin-dev/gallery/app/domain/file"
)

type FileRestService struct {
	repo     d.FileRepository
	s3client *s3.Client
}

type FileList struct {
	files []*FileDto
	total int
}

func NewFileRestService(fileRepository d.FileRepository, s3client *s3.Client) *FileRestService {
	return &FileRestService{fileRepository, s3client}
}

func (s *FileRestService) UploadFile(f io.Reader, name string) (*FileDto, error) {
	// я куда-то загрузил контент...
	file := d.File{
		Name: name,
	}

	s.repo.Create(&file)

	fileKey := strconv.Itoa(int(file.ID))
	bucketName := "gallery"

	_, err := s.s3client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
		Body:   f,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file %s to AWS S3: %e", name, err)
	}

	return NewFileDto(file), nil
}

func (s *FileRestService) getList() (*FileList, error) {
	files, err := s.repo.FindBy()
	if err != nil {
		return nil, err
	}

	dtos := make([]*FileDto, len(files))
	for i, file := range files {
		dtos[i] = NewFileDto(file)
	}

	return &FileList{dtos, 10}, nil
}
