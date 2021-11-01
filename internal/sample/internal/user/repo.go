package user

import "context"

type Repo interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserList(ctx context.Context, page, perpage int) (data []User, total int, err error)
}
