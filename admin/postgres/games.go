package postgres

import (
	"context"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/db"
	"github.com/resrrdttrt/VOU/pkg/errors"
	log "github.com/resrrdttrt/VOU/pkg/logger"
)

var _ admin.GameRepository = (*gamesRepository)(nil)

type gamesRepository struct {
	db db.Database
	l  log.Logger
}

func NewGameRepository(db db.Database, l log.Logger) admin.GameRepository {
	return &gamesRepository{
		db: db,
		l:  l,
	}
}

func (r *gamesRepository) GetAllGames(ctx context.Context) ([]admin.Game, error) {
	query := `SELECT * FROM games`
	params := map[string]interface{}{}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var games []admin.Game
	for rows.Next() {
		var game admin.Game
		if err := rows.StructScan(&game); err != nil {
			return nil, errors.Wrap(ErrSelectDb, err)
		}
		games = append(games, game)
	}
	return games, nil
}

func (r *gamesRepository) GetGameById(ctx context.Context, id string) (admin.Game, error) {
	query := `SELECT * FROM games WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := r.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return admin.Game{}, errors.Wrap(ErrSelectDb, err)
	}
	defer rows.Close()
	var game admin.Game
	if rows.Next() {
		if err := rows.StructScan(&game); err != nil {
			return admin.Game{}, errors.Wrap(ErrSelectDb, err)
		}
		return game, nil
	} else {
		return admin.Game{}, errors.Wrap(ErrNoData, err)
	}
}

func (r *gamesRepository) CreateGame(ctx context.Context, game admin.Game) error {
	query := `INSERT INTO games (name, images, type, exchange_allow, tutorial) VALUES (:name, :images, :type, :exchange_allow, :tutorial) RETURNING id`
	params := map[string]interface{}{
		"name":           game.Name,
		"images":         game.Images,
		"type":           game.Type,
		"exchange_allow": game.ExchangeAllow,
		"tutorial":       game.Tutorial,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrInsertDb, err)
	}
	return nil
}

func (r *gamesRepository) UpdateGame(ctx context.Context, game admin.Game) error {
	query := `UPDATE games SET name = :name, images = :images, type = :type, exchange_allow = :exchange_allow, tutorial = :tutorial WHERE id = :id`
	params := map[string]interface{}{
		"id":             game.ID,
		"name":           game.Name,
		"images":         game.Images,
		"type":           game.Type,
		"exchange_allow": game.ExchangeAllow,
		"tutorial":       game.Tutorial,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrUpdateDb, err)
	}
	return nil
}

func (r *gamesRepository) DeleteGame(ctx context.Context, id string) error {
	query := `DELETE FROM games WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return errors.Wrap(ErrDeleteDb, err)
	}
	return nil
}
