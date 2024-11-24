package services

import "github.com/DoCongThanhPhuong/go-backend/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetUserByID(id string) string {
	return us.userRepository.GetUserByID(id)
}