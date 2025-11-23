package repository

import (
    "database/sql"
    "mon-projet/internal/domain"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
    rows, err := r.DB.Query("SELECT id, email, name FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        var u domain.User
        if err := rows.Scan(&u.ID, &u.Email, &u.Name); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, rows.Err()
}

func (r *UserRepository) Create(u domain.User) error {
    _, err := r.DB.Exec("INSERT INTO users (email, name) VALUES ($1, $2)", u.Email, u.Name)
    return err
}
