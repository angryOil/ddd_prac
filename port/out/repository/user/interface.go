package user

type Repository interface {
	SearchUser(email, pw string) (int64, error)
	CreateUser(email, pw string) error
}
