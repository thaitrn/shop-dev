package service

import repo "shop-dev/internal/5.repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUser()
}
