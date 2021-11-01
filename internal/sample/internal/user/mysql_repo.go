package user

import (
	"context"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/jmoiron/sqlx"
)

type mysqlRepo struct {
	db     *sqlx.DB
	logger log.Logger
}

// NewMysqlRepo creates a new mysql user repo
func NewMysqlRepo(db *sqlx.DB, logger log.Logger) Repo {
	return mysqlRepo{db: db, logger: logger}
}

// GetUserList get user list
func (r mysqlRepo) GetUserList(ctx context.Context, page, perpage int) (data []User, total int, err error) {
	users := make([]User, 0)
	offset := (page - 1) * perpage
	query := "SELECT id, username, email, created_at, updated_at FROM users limit ?,?"

	err = r.db.Select(&users, query, offset, perpage)
	if err != nil {
		r.logger.Errorf("failed to get list of users %s", err)
		return
	}

	return users, len(users), nil
}

// GetUserByID get user by id
func (r mysqlRepo) GetUserByID(ctx context.Context, id string) (*User, error) {
	user := User{}
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	err := r.db.Get(&user, query, id)
	if err != nil {
		r.logger.Errorf("failed to get user by id %s, because of %s", id, err)
		return nil, err
	}

	return &user, nil
}
