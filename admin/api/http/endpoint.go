package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/common"
)

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
