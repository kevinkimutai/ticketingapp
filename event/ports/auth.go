package ports

type AuthPort interface {
	Verify(token string) (uint64, error)
}
