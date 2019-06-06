package interfaces

type TokenProvider interface {
	NewToken() (string, error)
	ValidateToken(token string) bool
}
