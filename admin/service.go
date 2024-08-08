package admin

import (
	"context"
	"time"

	log "github.com/resrrdttrt/VOU/pkg/logger"
)

type adminService struct {
	log       log.Logger
	users     UserRepository
	games     GameRepository
	statistic StatisticRepository
}

type Service interface {
	userService
	gameService
	statisticService
}

type userService interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserById(ctx context.Context, id string) (User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id string) error
	ActiveUser(ctx context.Context, id string) error
	DeactiveUser(ctx context.Context, id string) error
}

type gameService interface {
	GetAllGames(ctx context.Context) ([]Game, error)
	GetGameById(ctx context.Context, id string) (Game, error)
	CreateGame(ctx context.Context, game Game) error
	UpdateGame(ctx context.Context, game Game) error
	DeleteGame(ctx context.Context, id string) error
}

type statisticService interface {
	// General
	GetTotalUsers(ctx context.Context) (int, error)
	GetTotalGames(ctx context.Context) (int, error)

	GetTotalEnterprises(ctx context.Context) (int, error)
	GetTotalEndUser(ctx context.Context) (int, error)

	GetTotalActiveEndUsers(ctx context.Context) (int, error)
	GetTotalActiveEnterprises(ctx context.Context) (int, error)

	GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) (Statistic, error)
	GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) (Statistic, error)

	GetTotalNewEndUsersInWeek(ctx context.Context) (Statistic, error)
	GetTotalNewEnterprisesInWeek(ctx context.Context) (Statistic, error)
}

func NewAdminService(log log.Logger, users UserRepository, games GameRepository, statistic StatisticRepository) Service {
	return &adminService{
		log:       log,
		users:     users,
		games:     games,
		statistic: statistic,
	}
}

func (s *adminService) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.users.GetAllUsers(ctx)
}

func (s *adminService) GetUserById(ctx context.Context, id string) (User, error) {
	return s.users.GetUserById(ctx, id)
}

func (s *adminService) CreateUser(ctx context.Context, user User) error {
	return s.users.CreateUser(ctx, user)
}

func (s *adminService) UpdateUser(ctx context.Context, user User) error {
	return s.users.UpdateUser(ctx, user)
}

func (s *adminService) DeleteUser(ctx context.Context, id string) error {
	return s.users.DeleteUser(ctx, id)
}

func (s *adminService) ActiveUser(ctx context.Context, id string) error {
	user := User{
		ID:     id,
		Status: "active",
	}
	return s.users.UpdateUser(ctx, user)
}

func (s *adminService) DeactiveUser(ctx context.Context, id string) error {
	user := User{
		ID:     id,
		Status: "inactive",
	}
	return s.users.UpdateUser(ctx, user)
}

func (s *adminService) GetAllGames(ctx context.Context) ([]Game, error) {
	return s.games.GetAllGames(ctx)
}

func (s *adminService) GetGameById(ctx context.Context, id string) (Game, error) {
	return s.games.GetGameById(ctx, id)
}

func (s *adminService) CreateGame(ctx context.Context, game Game) error {
	return s.games.CreateGame(ctx, game)
}

func (s *adminService) UpdateGame(ctx context.Context, game Game) error {
	return s.games.UpdateGame(ctx, game)
}

func (s *adminService) DeleteGame(ctx context.Context, id string) error {
	return s.games.DeleteGame(ctx, id)
}

func (s *adminService) GetTotalUsers(ctx context.Context) (int, error) {
	return s.statistic.GetTotalUsers(ctx)
}

func (s *adminService) GetTotalGames(ctx context.Context) (int, error) {
	return s.statistic.GetTotalGames(ctx)
}

func (s *adminService) GetTotalEnterprises(ctx context.Context) (int, error) {
	return s.statistic.GetTotalEnterprises(ctx)
}

func (s *adminService) GetTotalEndUser(ctx context.Context) (int, error) {
	return s.statistic.GetTotalEndUser(ctx)
}

func (s *adminService) GetTotalActiveEndUsers(ctx context.Context) (int, error) {
	return s.statistic.GetTotalActiveEndUsers(ctx)
}

func (s *adminService) GetTotalActiveEnterprises(ctx context.Context) (int, error) {
	return s.statistic.GetTotalActiveEnterprises(ctx)
}

func (s *adminService) GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) (Statistic, error) {
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) (Statistic, error) {
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInWeek(ctx context.Context) (Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, now)
}

func (s *adminService) GetTotalNewEnterprisesInWeek(ctx context.Context) (Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, now)
}
