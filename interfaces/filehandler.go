package interfaces

type FileHandler interface {
	Read(file string) (string, error)
}
