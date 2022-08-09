package file

type FileRepository interface {
	Create(*File) error
	FindBy() ([]File, error)
	CountBy() (int64, error)
}
