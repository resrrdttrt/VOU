package admin

import (
	"context"
	"time"

	log "github.com/resrrdttrt/VOU/pkg/logger"
)

type adminService struct {
	log        log.Logger
	users      UserRepository
	games      GameRepository
	statistic  StatisticRepository
	auth       AuthRepository
	enterprise EnterpriseRepository
	event      EventRepository
	voucher    VoucherRepository
}

type Service interface {
	userService
	gameService
	statisticService
	authService
	enterpriseService
	eventService
	voucherService
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

	GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)
	GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error)

	GetTotalNewEndUsersInWeek(ctx context.Context) ([]Statistic, error)
	GetTotalNewEnterprisesInWeek(ctx context.Context) ([]Statistic, error)
}

type authService interface {
	// Auth
	Login(ctx context.Context, username, password string) (Token, error)
	GetUserIDByAccessToken(accessToken string) (string, error)
	GetUserRoleByID(userID string) (string, error)
}

type enterpriseService interface {
	RegisterEnterprise(ctx context.Context, enterprise Enterprise) error
	GetEnterpriseInfo(ctx context.Context) (Enterprise, error)
	UpdateEnterpriseInfo(ctx context.Context, enterprise Enterprise) error
}

type eventService interface {
	GetAllEvents(ctx context.Context) ([]Event, error)
	GetEventByID(ctx context.Context, id string) (Event, error)
	GetEventByTime(ctx context.Context, start time.Time, end time.Time) ([]Event, error)
	CreateEvent(ctx context.Context, event Event) error
	UpdateEvent(ctx context.Context, event Event) error
}

type voucherService interface {
	// GetAllVouchers(ctx context.Context) ([]Voucher, error)
	GetAllVouchersByEventID(ctx context.Context, eventID string) ([]Voucher, error)
	GetVoucherByID(ctx context.Context, id string, eventID string) (Voucher, error)
	CreateVoucher(ctx context.Context, voucher Voucher) error
	UpdateVoucher(ctx context.Context, voucher Voucher) error
	DeleteVoucher(ctx context.Context, id string, eventID string) error
}

func NewAdminService(log log.Logger, users UserRepository, games GameRepository, statistic StatisticRepository, auth AuthRepository, enterprise EnterpriseRepository, event EventRepository, voucher VoucherRepository) Service {
	return &adminService{
		log:       log,
		users:     users,
		games:     games,
		statistic: statistic,
		auth:      auth,
		enterprise: enterprise,
		event:     event,
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

func (s *adminService) GetTotalNewEnterprisesInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error) {
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInTime(ctx context.Context, start time.Time, end time.Time) ([]Statistic, error) {
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, end)
}

func (s *adminService) GetTotalNewEndUsersInWeek(ctx context.Context) ([]Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEndUsersInTime(ctx, start, now)
}

func (s *adminService) GetTotalNewEnterprisesInWeek(ctx context.Context) ([]Statistic, error) {
	now := time.Now()
	start := now.AddDate(0, 0, -7)
	return s.statistic.GetTotalNewEnterprisesInTime(ctx, start, now)
}

func (s *adminService) Login(ctx context.Context, username, password string) (Token, error) {
	return s.auth.Login(ctx, username, password)
}

func (s *adminService) GetUserIDByAccessToken(accessToken string) (string, error) {
	return s.auth.GetUserIDByAccessToken(accessToken)
}

func (s *adminService) GetUserRoleByID(userID string) (string, error) {
	return s.auth.GetUserRoleByID(userID)
}

func (s *adminService) RegisterEnterprise(ctx context.Context, enterprise Enterprise) error {
	return s.enterprise.CreateEnterprise(ctx, enterprise)
}

func (s *adminService) GetEnterpriseInfo(ctx context.Context) (Enterprise, error) {
	enterpriseID := ctx.Value("userID").(string)
	return s.enterprise.GetEnterpriseByID(ctx, enterpriseID)
}

func (s *adminService) UpdateEnterpriseInfo(ctx context.Context, enterprise Enterprise) error {
	return s.enterprise.UpdateEnterprise(ctx, enterprise)
}

func (s *adminService) GetAllEvents(ctx context.Context) ([]Event, error) {
	enterpriseID := ctx.Value("userID").(string)
	return s.event.GetAllEventsByEnterpriseID(ctx, enterpriseID)
}

func (s *adminService) GetEventByID(ctx context.Context, id string) (Event, error) {
	enterpriseID := ctx.Value("userID").(string)
	return s.event.GetEventByID(ctx, id, enterpriseID)
}

func (s *adminService) GetEventByTime(ctx context.Context, start time.Time, end time.Time) ([]Event, error) {
	enterpriseID := ctx.Value("userID").(string)
	return s.event.GetEventByTime(ctx,enterpriseID, start, end)
}

func (s *adminService) CreateEvent(ctx context.Context, event Event) error {
	return s.event.CreateEvent(ctx, event)
}

func (s *adminService) UpdateEvent(ctx context.Context, event Event) error {
	return s.event.UpdateEvent(ctx, event)
}

// func (s *adminService) GetAllVouchers(ctx context.Context) ([]Voucher, error) {
// 	return s.voucher.GetAllVouchers(ctx)
// }

func (s *adminService) GetAllVouchersByEventID(ctx context.Context, eventID string) ([]Voucher, error) {
	return s.voucher.GetAllVouchersByEventID(ctx, eventID)
}

func (s *adminService) GetVoucherByID(ctx context.Context, id string, eventID string) (Voucher, error) {
	return s.voucher.GetVoucherByID(ctx, id, eventID)
}

func (s *adminService) CreateVoucher(ctx context.Context, voucher Voucher) error {
	return s.voucher.CreateVoucher(ctx, voucher)
}

func (s *adminService) UpdateVoucher(ctx context.Context, voucher Voucher) error {
	return s.voucher.UpdateVoucher(ctx, voucher)
}

func (s *adminService) DeleteVoucher(ctx context.Context, id string, eventID string) error {
	return s.voucher.DeleteVoucher(ctx, id, eventID)
}


