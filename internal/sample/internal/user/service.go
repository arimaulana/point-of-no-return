package user

import (
	"context"
	"time"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/generator"
	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
)

type Service interface {
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserList(ctx context.Context, page, perpage int) (data []User, total int, err error)

	// AddUser(ctx context.Context, username, email, password string) (User, error)
	// UpdateUser(ctx context.Context, id, username, email, password string) (User, error)
	// RemoveUser(ctx context.Context, id string) (User, error)
}

type userService struct {
	repo   Repo
	logger log.Logger
}

// NewService return new userService instance
func NewService(repo Repo, logger log.Logger) Service {
	return userService{repo: repo, logger: logger}
}

func (s userService) GetUserByID(ctx context.Context, id string) (User, error) {
	return User{
		ID:        generator.GenerateUUID(),
		UserName:  "admin",
		Email:     "admin@sample.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
	// return s.repo.GetUserByID(ctx, id)
}

func (s userService) GetUserList(ctx context.Context, page, perpage int) (data []User, total int, err error) {

	// data = s.repo.GetUserList(ctx, id)
	admin := User{
		ID:        generator.GenerateUUID(),
		UserName:  "admin",
		Email:     "admin@sample.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data = []User{
		admin,
	}
	total = len(data)

	return
}
