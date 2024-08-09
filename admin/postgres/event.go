package postgres

import (
	"context"
	"time"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.EventRepository = (*eventRepository)(nil)

type eventRepository struct {
	db db.Database
	l  log.Logger
}

func NewEventRepository(db db.Database, l log.Logger) admin.EventRepository {
	return &eventRepository{
		db: db,
		l:  l,
	}
}


func (r *eventRepository) GetEventByID(ctx context.Context, id string, enterprise_id string) (admin.Event, error) {
	query := `SELECT * FROM events WHERE id = :id AND user_id = :user_id`
	params := map[string]interface{}{
		"id":           id,
		"user_id": enterprise_id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.Event{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var event admin.Event
	if rows.Next() {
		if err := rows.StructScan(&event); err != nil {
			return admin.Event{}, errors.Wrap(ErrSelectDb, err)
		}
		return event, nil
	} else {
		return admin.Event{}, errors.Wrap(ErrNoData, err)
	}
}


func (r *eventRepository) CreateEvent(ctx context.Context, event admin.Event) error {
	query := `INSERT INTO events (id, name, images, voucher_num, start_time, end_time, game_id, user_id) VALUES (:id, :name, :images, :voucher_num, :start_time, :end_time, :game_id, :user_id) RETURNING id`
	params := map[string]interface{}{
		"name":        event.Name,
		"images":      event.Images,
		"voucher_num": event.VoucherNum,
		"start_time":  event.StartTime,
		"end_time":    event.EndTime,
		"game_id":     event.GameID,
		"user_id":     event.UserID,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, event admin.Event) error {
	query := `UPDATE events SET `
	params := map[string]interface{}{
		"id": event.ID,
		"user_id": event.UserID,
	}

	if event.Name != "" {
		query += `name = :name, `
		params["name"] = event.Name
	}

	if event.Images != "" {
		query += `images = :images, `
		params["images"] = event.Images
	}

	if event.VoucherNum != 0 {
		query += `voucher_num = :voucher_num, `
		params["voucher_num"] = event.VoucherNum
	}

	if !event.StartTime.IsZero() {
		query += `start_time = :start_time, `
		params["start_time"] = event.StartTime
	}

	if !event.EndTime.IsZero() {
		query += `end_time = :end_time, `
		params["end_time"] = event.EndTime
	}

	if event.GameID != "" {
		query += `game_id = :game_id, `
		params["game_id"] = event.GameID
	}

	query = query[:len(query)-2] + ` WHERE id = :id and user_id = :user_id RETURNING *`
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}

func (r *eventRepository) GetAllEventsByEnterpriseID(ctx context.Context, enterprise_id string) ([]admin.Event, error) {
	query := `SELECT * FROM events WHERE user_id = :user_id`
	params := map[string]interface{}{
		"user_id": enterprise_id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var events []admin.Event
	for rows.Next() {
		var event admin.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, errors.Wrap(ErrSelectDb, err)
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventRepository) GetEventByTime(ctx context.Context, enterprise_id string, start time.Time, end time.Time) ([]admin.Event, error) {
	query := `SELECT * FROM events WHERE user_id = :user_id AND start_time >= :start_time AND end_time <= :end_time`
	params := map[string]interface{}{
		"user_id":    enterprise_id,
		"start_time": start,
		"end_time":   end,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var events []admin.Event
	for rows.Next() {
		var event admin.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, errors.Wrap(ErrSelectDb, err)
		}
		events = append(events, event)
	}
	return events, nil
}