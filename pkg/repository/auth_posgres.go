package repository

import (
	"app"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password) values ($1, $2) RETURNING id", userTable)
	user_row := r.db.QueryRow(query, user.Email, user.Password)
	if err := user_row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

// func (r *AuthPostgres) GetUserById(id int) (app.User, error) {
// 	var user app.User
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", userTable)
// 	err := r.db.Get(&user, query, id)
// 	return user, err
// }
