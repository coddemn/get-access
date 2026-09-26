package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/coddemn/get-access/internal/api/apperrors"
	"github.com/coddemn/get-access/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetById(id int) (domain.User, error) {
	var user domain.User

	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("GetUserByID %d: unknown user", id)
		}
		return user, fmt.Errorf("GetUserByID %d: %v", id, err)
	}

	return user, nil
}

func (r *UserRepository) GetByLogin(login, pass string) (domain.User, error) {

	var user domain.User

	err := r.db.QueryRow("SELECT * FROM users WHERE login = $1 and pass = $2", login, pass).Scan(&user.ID, &user.Name, &user.Login, &user.Pass)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Auth err: %v", err)
			return user, fmt.Errorf("Auth - incorrect data: %w", apperrors.ErrIncorrectAuth)
		}
		return user, fmt.Errorf("Auth: %v", err)
	}

	return user, nil
}

func (r *UserRepository) Add(userData domain.User) (int64, error) {

	row := r.db.QueryRow("INSERT INTO users (name, login, pass) VALUES ($1, $2, $3) RETURNING id", userData.Name, userData.Login, userData.Pass)

	var id int64
	if err := row.Scan(&id); err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("failed to create user: %w", apperrors.ErrNameTaken)
		}

		return 0, fmt.Errorf("Registrate: %v", err)
	}

	return id, nil
}
