package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/middlewares"
	"github.com/resrrdttrt/VOU/pkg/errors"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
)

func MakeHandler(svc admin.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := bone.New()

	r.Get("/admin/user", kithttp.NewServer(
		getAllUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/user/:id", kithttp.NewServer(
		getUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/admin/user", kithttp.NewServer(
		createUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
		opts...,
	))
	r.Put("/admin/user/:id", kithttp.NewServer(
		updateUserEndpoint(svc),
		decodeUpdateUserRequest,
		encodeResponse,
		opts...,
	))
	r.Delete("/admin/user/:id", kithttp.NewServer(
		deleteUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/user/active/:id", kithttp.NewServer(
		activeUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/user/deactive/:id", kithttp.NewServer(
		deactiveUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/game", kithttp.NewServer(
		getAllGamesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/game/:id", kithttp.NewServer(
		getGameEndpoint(svc),
		decodeGetGameRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/admin/game", kithttp.NewServer(
		createGameEndpoint(svc),
		decodeCreateGameRequest,
		encodeResponse,
		opts...,
	))
	r.Put("/admin/game/:id", kithttp.NewServer(
		updateGameEndpoint(svc),
		decodeUpdateGameRequest,
		encodeResponse,
		opts...,
	))
	r.Delete("/admin/game/:id", kithttp.NewServer(
		deleteGameEndpoint(svc),
		decodeGetGameRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_users", kithttp.NewServer(
		getTotalUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_games", kithttp.NewServer(
		getTotalGamesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_enterprises", kithttp.NewServer(
		getTotalEnterprisesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_end_users", kithttp.NewServer(
		getTotalEndUserEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_active_end_users", kithttp.NewServer(
		getTotalActiveEndUsersEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_active_enterprises", kithttp.NewServer(
		getTotalActiveEnterprisesEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_new_enterprises_in_time", kithttp.NewServer(
		getTotalNewEnterprisesInTimeEndpoint(svc),
		decodeStatisticInTimeRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_new_end_users_in_time", kithttp.NewServer(
		getTotalNewEndUsersInTimeEndpoint(svc),
		decodeStatisticInTimeRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_new_end_users_in_week", kithttp.NewServer(
		getTotalNewEndUsersInWeekEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/admin/statistic/total_new_enterprises_in_week", kithttp.NewServer(
		getTotalNewEnterprisesInWeekEndpoint(svc),
		decodeNothingRequest,
		encodeResponse,
		opts...,
	))
	handler := middlewares.VerifyAdminMiddleware(r)
	return handler
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if ar, ok := response.(Response); ok {
		fmt.Println(ar)
		for k, v := range ar.Headers() {
			w.Header().Set(k, v)
		}
		w.WriteHeader(ar.Code())
		if ar.Empty() {
			return nil
		}
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch errorVal := err.(type) {
	case errors.Error:
		switch {
		case errors.Contains(errorVal, errors.ErrNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Contains(errorVal, errors.ErrUnsupportedMediaType):
			w.WriteHeader(http.StatusUnsupportedMediaType)
		case errors.Contains(errorVal, errors.ErrMalformedEntity):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Contains(errorVal, errors.ErrBadRequest):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		if errorVal.Msg() != "" {
			if err := json.NewEncoder(w).Encode(errorResponse{Message: errorVal.Msg(), Code: errorVal.Code(), Error: errorVal.Error()}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getUserRequest
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeGetGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getGameRequest
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeCreateGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}

func decodeUpdateGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	id := bone.GetValue(r, "id")
	req.ID = id
	return req, nil
}

func decodeNothingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeStatisticInTimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req statisticInTimeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
	}
	return req, nil
}


