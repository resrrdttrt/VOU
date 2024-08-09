package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/common"
)

func getAllUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetAllUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(users), nil
	}
}

func createUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user := admin.User{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
			Role:     req.Role,
			Status:   req.Status,
		}
		if err := svc.CreateUser(ctx, user); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func updateUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user := admin.User{
			ID:       req.ID,
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
			Role:     req.Role,
			Status:   req.Status,
		}
		if err := svc.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		user, err := svc.GetUserById(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(user), nil
	}
}

func deleteUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeleteUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func activeUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.ActiveUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func deactiveUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeactiveUser(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getAllGamesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		games, err := svc.GetAllGames(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(games), nil
	}
}

func getGameEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getGameRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		game, err := svc.GetGameById(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(game), nil
	}
}

func createGameEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createGameRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		game := admin.Game{
			Name:          req.Name,
			Images:        req.Images,
			Type:          req.Type,
			ExchangeAllow: req.ExchangeAllow,
			Tutorial:      req.Tutorial,
		}
		if err := svc.CreateGame(ctx, game); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func updateGameEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateGameRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		game := admin.Game{
			ID:            req.ID,
			Name:          req.Name,
			Images:        req.Images,
			Type:          req.Type,
			ExchangeAllow: req.ExchangeAllow,
			Tutorial:      req.Tutorial,
		}
		if err := svc.UpdateGame(ctx, game); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func deleteGameEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getGameRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeleteGame(ctx, req.ID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getTotalUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := svc.GetTotalUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(users), nil
	}
}

func getTotalGamesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		games, err := svc.GetTotalGames(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(games), nil
	}
}

func getTotalEnterprisesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalEnterprises(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalEndUserEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalEndUser(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalActiveEndUsersEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalActiveEndUsers(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalActiveEnterprisesEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalActiveEnterprises(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalNewEnterprisesInTimeEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req statisticInTimeRequest
		if err := req.validate(); err != nil {
			return nil, err
		}
		enterprises, err := svc.GetTotalNewEnterprisesInTime(ctx, req.Start, req.End)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func getTotalNewEndUsersInTimeEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req statisticInTimeRequest
		if err := req.validate(); err != nil {
			return nil, err
		}
		endusers, err := svc.GetTotalNewEndUsersInTime(ctx, req.Start, req.End)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalNewEndUsersInWeekEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		endusers, err := svc.GetTotalNewEndUsersInWeek(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(endusers), nil
	}
}

func getTotalNewEnterprisesInWeekEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprises, err := svc.GetTotalNewEnterprisesInWeek(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprises), nil
	}
}

func loginEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		token, err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(token), nil
	}
}

func registerEnterpriseEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(enterpriseRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		enterprise := admin.Enterprise{
			Name:     req.Name,
			Field:    req.Field,
			Location: req.Location,
			GPS:      req.GPS,
			Status:   req.Status,
		}
		if err := svc.RegisterEnterprise(ctx, enterprise); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getEnterpriseInfoEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		enterprise, err := svc.GetEnterpriseInfo(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(enterprise), nil
	}
}

func updateEnterpriseInfoEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(enterpriseRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		enterprise := admin.Enterprise{
			Name:     req.Name,
			Field:    req.Field,
			Location: req.Location,
			GPS:      req.GPS,
			Status:   req.Status,
		}
		if err := svc.UpdateEnterpriseInfo(ctx, enterprise); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getAllEventsEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		events, err := svc.GetAllEvents(ctx)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(events), nil
	}
}

func getEventByIDEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEventIDRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		event, err := svc.GetEventByID(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(event), nil
	}
}


func createEventEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createEventRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		event := admin.Event{
			Name:       req.Name,
			Images:     req.Images,
			VoucherNum: req.VoucherNum,
			StartTime:  req.StartTime,
			EndTime:    req.EndTime,
			GameID:     req.GameID,
			UserID:     req.UserID,
		}
		if err := svc.CreateEvent(ctx, event); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func updateEventEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEventRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		event := admin.Event{
			ID:         req.ID,
			Name:       req.Name,
			Images:     req.Images,
			VoucherNum: req.VoucherNum,
			StartTime:  req.StartTime,
			EndTime:    req.EndTime,
			GameID:     req.GameID,
			UserID:     req.UserID,
		}
		if err := svc.UpdateEvent(ctx, event); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func getAllVouchersByEventIDEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEventIDRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		vouchers, err := svc.GetAllVouchersByEventID(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(vouchers), nil
	}
}

func getVoucherByIDEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getVoucherByIDRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		voucher, err := svc.GetVoucherByID(ctx, req.ID, req.EventID)
		if err != nil {
			return nil, err
		}
		return common.SuccessRes(voucher), nil
	}
}

func createVoucherEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createVoucherRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		voucher := admin.Voucher{
			Code:        req.Code,
			Qrcode:      req.Qrcode,
			Images:      req.Images,
			Value:       req.Value,
			Description: req.Description,
			ExpiredTime: req.ExpiredTime,
			Status:      req.Status,
			EventID:     req.EventID,
		}
		if err := svc.CreateVoucher(ctx, voucher); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func updateVoucherEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateVoucherRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		voucher := admin.Voucher{
			ID:          req.ID,
			Code:        req.Code,
			Qrcode:      req.Qrcode,
			Images:      req.Images,
			Value:       req.Value,
			Description: req.Description,
			ExpiredTime: req.ExpiredTime,
			Status:      req.Status,
			EventID:     req.EventID,
		}
		if err := svc.UpdateVoucher(ctx, voucher); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}

func deleteVoucherEndpoint(svc admin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getVoucherByIDRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if err := svc.DeleteVoucher(ctx, req.ID, req.EventID); err != nil {
			return nil, err
		}
		return common.SuccessRes(nil), nil
	}
}