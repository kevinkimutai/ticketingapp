package ports

type DBPort interface {
	CreateUser()
	GetUserByEmail()
}
