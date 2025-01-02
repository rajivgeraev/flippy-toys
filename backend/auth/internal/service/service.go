package service

import "github.com/rajivgeraev/flippy-toys/backend/auth/internal/model"

type UserService interface {
    CreateUser(user *model.User) error
    GetUser(telegramID int64) (*model.User, error)
    UpdateUser(user *model.User) error
}

type Services struct {
    User UserService
}

type Deps struct {
    UserRepo repository.UserRepository
}

func NewServices(deps Deps) *Services {
    return &Services{
        User: NewUserService(deps.UserRepo),
    }
}