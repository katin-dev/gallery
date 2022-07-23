package file

type FileRepository interface {
	Create(*File)
	FindBy() []File
}
