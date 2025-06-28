package sqlite

import (
	"context"
	"database/sql"

	"github/k-tsurumaki/quilldeck/internal/domain/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *DB
}

func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID.String(),
		user.Email,
		user.Password,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT id, email, password, name, created_at, updated_at FROM users WHERE id = ?`

	var user models.User
	var idStr string
	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&idStr,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user.ID = uuid.MustParse(idStr)
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password, name, created_at, updated_at FROM users WHERE email = ?`

	var user models.User
	var idStr string
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&idStr,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user.ID = uuid.MustParse(idStr)
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET email = ?, password = ?, name = ?, updated_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Password,
		user.Name,
		user.UpdatedAt,
		user.ID.String(),
	)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}