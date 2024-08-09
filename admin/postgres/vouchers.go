package postgres

import (
	"context"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.VoucherRepository = (*voucherRepository)(nil)

type voucherRepository struct {
	db db.Database
	l  log.Logger
}

func NewVoucherRepository(db db.Database, l log.Logger) admin.VoucherRepository {
	return &voucherRepository{
		db: db,
		l:  l,
	}
}

func (r *voucherRepository) CreateVoucher(ctx context.Context, voucher admin.Voucher) error {
	query := `INSERT INTO vouchers (code, qrcode, images, value, description, expired_time, status, event_id) VALUES (:code, :qrcode, :images, :value, :description, :expired_time, :status, :event_id) RETURNING id`
	params := map[string]interface{}{
		"code":         voucher.Code,
		"qrcode":       voucher.Qrcode,
		"images":       voucher.Images,
		"value":        voucher.Value,
		"description":  voucher.Description,
		"expired_time": voucher.ExpiredTime,
		"status":       voucher.Status,
		"event_id":     voucher.EventID,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}


// func (r *voucherRepository) GetAllVouchers(ctx context.Context) ([]admin.Voucher, error) {
// 	query := `SELECT * FROM vouchers`
// 	params := map[string]interface{}{}
// 	rows, err := r.db.NamedQueryContext(ctx, query, params)
// 	if err != nil {
// 		return nil, errors.Wrap(ErrSelectDb, err)
// 	}
// 	defer rows.Close()
// 	var vouchers []admin.Voucher
// 	for rows.Next() {
// 		var voucher admin.Voucher
// 		if err := rows.StructScan(&voucher); err != nil {
// 			return nil, errors.Wrap(ErrSelectDb, err)
// 		}
// 		vouchers = append(vouchers, voucher)
// 	}
// 	return vouchers, nil

// 	// TODO: WRONG IMPLEMENTATION
// }

func (r *voucherRepository) GetAllVouchersByEventID(ctx context.Context, eventID string) ([]admin.Voucher, error) {
	query := `SELECT * FROM vouchers WHERE event_id = :event_id`
	params := map[string]interface{}{
		"event_id": eventID,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var vouchers []admin.Voucher
	for rows.Next() {
		var voucher admin.Voucher
		if err := rows.StructScan(&voucher); err != nil {
			return nil, errors.Wrap(ErrSelectDb, err)
		}
		vouchers = append(vouchers, voucher)
	}
	return vouchers, nil
}

func (r *voucherRepository) GetVoucherByID(ctx context.Context, id string, eventID string) (admin.Voucher, error) {
	query := `SELECT * FROM vouchers WHERE id = :id AND event_id = :event_id`
	params := map[string]interface{}{
		"id":       id,
		"event_id": eventID,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.Voucher{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var voucher admin.Voucher
	if rows.Next() {
		if err := rows.StructScan(&voucher); err != nil {
			return admin.Voucher{}, errors.Wrap(ErrSelectDb, err)
		}
		return voucher, nil
	} else {
		return admin.Voucher{}, errors.Wrap(ErrNoData, err)
	}
}

func (r *voucherRepository) UpdateVoucher(ctx context.Context, voucher admin.Voucher) error {
	query := `UPDATE vouchers SET `
	params := map[string]interface{}{
		"id": voucher.ID,
		"event_id": voucher.EventID,
	}

	if voucher.Code != "" {
		query += `code = :code, `
		params["code"] = voucher.Code
	}

	if voucher.Qrcode != "" {
		query += `qrcode = :qrcode, `
		params["qrcode"] = voucher.Qrcode
	}

	if voucher.Images != "" {
		query += `images = :images, `
		params["images"] = voucher.Images
	}

	if voucher.Value != 0 {
		query += `value = :value, `
		params["value"] = voucher.Value
	}

	if voucher.Description != "" {
		query += `description = :description, `
		params["description"] = voucher.Description
	}

	if !voucher.ExpiredTime.IsZero() {
		query += `expired_time = :expired_time, `
		params["expired_time"] = voucher.ExpiredTime
	}

	if voucher.Status != "" {
		query += `status = :status, `
		params["status"] = voucher.Status
	}

	query = query[:len(query)-2] + ` WHERE id = :id and event_id = :event_id RETURNING *`
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}

func (r *voucherRepository) DeleteVoucher(ctx context.Context, id string, eventID string) error {
	query := `DELETE FROM vouchers WHERE id = :id AND event_id = :event_id`
	params := map[string]interface{}{
		"id":       id,
		"event_id": eventID,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrDeleteDb, err)
	}
	return nil
}
