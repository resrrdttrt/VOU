package postgres

import (
	"context"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.EnterpriseRepository = (*enterpriseRepository)(nil)

type enterpriseRepository struct {
	db db.Database
	l  log.Logger
}

func NewEnterpriseRepository(db db.Database, l log.Logger) admin.EnterpriseRepository {
	return &enterpriseRepository{
		db: db,
		l:  l,
	}
}

func (r *enterpriseRepository) CreateEnterprise(ctx context.Context, enterprise admin.Enterprise) error {
	query := `INSERT INTO enterprises (id, name, field, location, gps, status, created_at, updated_at) VALUES (:id, :name, :field, :location, :gps, :status) RETURNING id`
	params := map[string]interface{}{
		"name":     enterprise.Name,
		"field":    enterprise.Field,
		"location": enterprise.Location,
		"gps":      enterprise.GPS,
		"status":   enterprise.Status,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}

func (r *enterpriseRepository) GetEnterpriseByID(ctx context.Context, id string) (admin.Enterprise, error) {
	var enterprise admin.Enterprise
	query := `SELECT * FROM enterprises WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.Enterprise{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&enterprise); err != nil {
			return admin.Enterprise{}, errors.Wrap(ErrSelectDb, err)
		}
		return enterprise, nil
	} else {
		return admin.Enterprise{}, errors.Wrap(ErrNoData, err)
	}
}

func (r *enterpriseRepository) UpdateEnterprise(ctx context.Context, enterprise admin.Enterprise) error {
	query := `UPDATE enterprises SET `
	params := map[string]interface{}{
		"id": enterprise.ID,
	}

	if enterprise.Name != "" {
		query += `name = :name, `
		params["name"] = enterprise.Name
	}

	if enterprise.Field != "" {
		query += `field = :field, `
		params["field"] = enterprise.Field
	}

	if enterprise.Location != "" {
		query += `location = :location, `
		params["location"] = enterprise.Location
	}

	if enterprise.GPS != "" {
		query += `gps = :gps, `
		params["gps"] = enterprise.GPS
	}

	if enterprise.Status != "" {
		query += `status = :status, `
		params["status"] = enterprise.Status
	}

	query = query[:len(query)-2] + ` WHERE id = :id RETURNING *`
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}
