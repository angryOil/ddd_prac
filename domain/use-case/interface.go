package use_case

type UserService interface {
	Login(email, pw string) (l *LoginUser, err error)
	Register(email, pw string) (err error)
}
