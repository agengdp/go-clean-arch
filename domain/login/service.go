package login

type Service interface {
	Login(*Login) error
}

type service struct{
	repo Repository
}

func NewService(repo Repository) Service{
	return &service{repo}
}

func (svc *service) Login(u *Login) error{
	return svc.repo.PerformLogin(u)
}