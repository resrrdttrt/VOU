package postgres

import (
	"context"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.UserRepository = (*usersRepository)(nil)

type usersRepository struct {
	db db.Database
	l  log.Logger
}

func NewUserRepository(db db.Database, l log.Logger) admin.UserRepository {
	return &usersRepository{
		db: db,
		l:  l,
	}
}

func (r *usersRepository) GetUserById(ctx context.Context, id string) (admin.User, error) {
	query := `SELECT * FROM users WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.User{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var user admin.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return admin.User{}, errors.Wrap(ErrSelectDb, err)
		}
		return user, nil
	} else {
		return admin.User{}, nil
	}
}

func (r *usersRepository) CreateUser(ctx context.Context, user admin.User) error {
	query := `INSERT INTO users (name, username, password, email, phone, role, status) VALUES (:name, :username, :password, :email, :phone, :role, :status) RETURNING id`
	params := map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"phone":    user.Phone,
		"role":     user.Role,
		"status":   user.Status,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}

func (r *usersRepository) UpdateUser(ctx context.Context, user admin.User) error {
	query := `UPDATE users SET name = :name, username = :username, password = :password, email = :email, phone = :phone, role = :role, status = :status WHERE id = :id RETURNING *`
	params := map[string]interface{}{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
		"phone":    user.Phone,
		"role":     user.Role,
		"status":   user.Status,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}

func (r *usersRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrDeleteDb, err)
	}
	return nil
}
