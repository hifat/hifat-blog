package authService

import (
	"github.com/hifat/hifat-blog-api/internal/domain"
	"github.com/hifat/hifat-blog-api/internal/utils"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{userRepo}
}

func (u authService) Register(req domain.PayloadUser) (res *domain.ResponseUser, err error) {
	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	res, err = u.userRepo.Create(req)

	return res, err
}
