package ports

type AuthPort interface {
	Verify(token string) error
}
