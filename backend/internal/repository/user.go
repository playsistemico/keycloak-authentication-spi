// backend/internal/repository/repository.go
package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"backend/internal/domain"
)

const (
	tableUsers = "users"
)

type UsersRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
}

type usersRepository struct {
	db         *sqlx.DB
	sqlBuilder *usersSQL
}

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{
		db:         db,
		sqlBuilder: &usersSQL{},
	}
}

func (r *usersRepository) Create(u *domain.User) (*domain.User, error) {
	query, args, err := r.sqlBuilder.CreateSQL(u)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *usersRepository) GetByID(id string) (*domain.User, error) {
	query, args, err := r.sqlBuilder.GetByIDSQL(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = r.db.QueryRowx(query, args...).StructScan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}
