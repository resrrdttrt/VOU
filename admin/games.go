package admin

import (
	"context"
	"time"
)

type Game struct {
	ID            string    `db:"id" json:"id,omitempty"`
	Name          string    `db:"name" json:"name"`
	Images        string    `db:"images" json:"images"`
	Type          string    `db:"type" json:"type"`
	CreatedAt     time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at,omitempty"`
	ExchangeAllow bool      `db:"exchange_allow" json:"exchange_allow"`
	Tutorial      string    `db:"tutorial" json:"tutorial"`
}

type GameRepository interface {
	GetAllGames(ctx context.Context) ([]Game, error)
	GetGameById(ctx context.Context, id string) (Game, error)
	CreateGame(ctx context.Context, game Game) error
	UpdateGame(ctx context.Context, game Game) error
	DeleteGame(ctx context.Context, id string) error
}
