package file

type FileRepository interface {
	FindBy() []File
}
