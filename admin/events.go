package admin

import (
	"context"
	"time"
)

type Event struct {
	ID         string    `db:"id" json:"id,omitempty"`
	Name       string    `db:"name" json:"name"`
	Images     string    `db:"images" json:"images"`
	VoucherNum int       `db:"voucher_num" json:"voucher_num"`
	StartTime  time.Time `db:"start_time" json:"start_time"`
	EndTime    time.Time `db:"end_time" json:"end_time"`
	GameID     string    `db:"game_id" json:"game_id"`
	UserID     string    `db:"user_id" json:"user_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type EventRepository interface {
	GetEventByID(ctx context.Context, id string, enterprise_id string) (Event, error)
	CreateEvent(ctx context.Context, event Event) error
	UpdateEvent(ctx context.Context, event Event) error
	GetAllEventsByEnterpriseID(ctx context.Context, enterprise_id string) ([]Event, error)
	GetEventByTime(ctx context.Context, enterprise_id string, start time.Time, end time.Time) ([]Event, error)
}
