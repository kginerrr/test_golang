package repository

import (
	"database/sql"
	"errors"
	"fibertesttask/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
