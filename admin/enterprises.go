package admin

import (
	"context"
	"time"
)

type Enterprise struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name"`
	Field     string    `db:"field" json:"field"`
	Location  string    `db:"location" json:"location"`
	GPS       string    `db:"gps" json:"gps"`
	Status    string    `db:"status" json:"status,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type EnterpriseRepository interface {
	GetEnterpriseByID(ctx context.Context, id string) (Enterprise, error)
	CreateEnterprise(ctx context.Context, enterprise Enterprise) error
	UpdateEnterprise(ctx context.Context, enterprise Enterprise) error
}
