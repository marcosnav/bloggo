package interfaces

type BcryptHasher interface {
	Generate(password string) (string, error)
	Compare(hash string, password string) error
}
