package admin

import (
	"context"
	"time"
)

type Game struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name"`
	Images        string    `json:"images"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ExchangeAllow bool      `json:"exchange_allow"`
	Tutorial      string    `json:"tutorial"`
}

type GameRepository interface {
	GetGameById(ctx context.Context, id string) (Game, error)
	CreateGame(ctx context.Context, game Game) error
	UpdateGame(ctx context.Context, game Game) error
	DeleteGame(ctx context.Context, id string) error
}
