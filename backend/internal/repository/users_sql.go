// backend/internal/repository/users_sql.go
package repository

import (
	"fmt"

	"backend/internal/domain"
)

type usersSQL struct{}

func (u *usersSQL) CreateSQL(user *domain.User) (string, []interface{}, error) {
	query := fmt.Sprintf("INSERT INTO %s (id, name) VALUES (?, ?)", tableUsers)
	args := []interface{}{user.ID, user.Name}
	return query, args, nil
}

func (u *usersSQL) GetByIDSQL(id string) (string, []interface{}, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = ?", tableUsers)
	args := []interface{}{id}
	return query, args, nil
}
