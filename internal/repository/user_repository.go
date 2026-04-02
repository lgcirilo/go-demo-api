package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lgcirilo/go-demo-api/internal/models"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	return users, nil
}
