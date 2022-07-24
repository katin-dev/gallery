package file

type FileRepository interface {
	Create(*File) error
	FindBy() ([]File, error)
}
