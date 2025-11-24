package db

import (
	"database/sql"
	"mon-projet/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, email, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) Create(u model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (email, name) VALUES ($1, $2)", u.Email, u.Name)
	return err
}

func (r *UserRepository) FindByID(id int64) (model.User, error) {
	var u model.User
	err := r.DB.QueryRow("SELECT id, email, name FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email, &u.Name)
	return u, err
}

func (r *UserRepository) Update(id int64, u model.User) error {
	res, err := r.DB.Exec("UPDATE users SET email = $1, name = $2 WHERE id = $3", u.Email, u.Name, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	res, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
