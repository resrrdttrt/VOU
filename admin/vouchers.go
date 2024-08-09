package admin

import (
	"context"
	"time"
)

type Voucher struct {
	ID          string    `db:"id" json:"id,omitempty"`
	Code        string    `db:"code" json:"code"`
	Qrcode      string    `db:"qrcode" json:"qrcode"`
	Images      string    `db:"images" json:"images"`
	Value       int       `db:"value" json:"value"`
	Description string    `db:"description" json:"description"`
	ExpiredTime time.Time `db:"expired_time" json:"expired_time"`
	Status      string    `db:"status" json:"status"`
	EventID     string    `db:"event_id" json:"event_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type VoucherRepository interface {
	// GetAllVouchers(ctx context.Context) ([]Voucher, error)
	GetAllVouchersByEventID(ctx context.Context, eventID string) ([]Voucher, error)
	GetVoucherByID(ctx context.Context, id string, eventID string) (Voucher, error)
	CreateVoucher(ctx context.Context, voucher Voucher) error
	UpdateVoucher(ctx context.Context, voucher Voucher) error
	DeleteVoucher(ctx context.Context, id string, eventID string) error
}
