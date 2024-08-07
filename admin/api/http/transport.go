package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/resrrdttrt/VOU/admin"
	"github.com/resrrdttrt/VOU/pkg/errors"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
)

func MakeHandler(svc admin.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := bone.New()

	r.Get("/admin/user/:id", kithttp.NewServer(
		getUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/admin/user/", kithttp.NewServer(
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
	// r.Post("/mysafe/feedback/:api", kithttp.NewServer(
	// 	feedbackEndpoint(svc),
	// 	decodeFeedbackRequest,
	// 	encodeResponse,
	// ))
	return r
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

// func decodeAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req authRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		return nil, errors.Wrap(errors.ErrMalformedEntity, err)
// 	}
// 	extraInfo, err := keto.NewExtra(r)
// 	if err != nil {
// 		return nil, errors.Wrap(keto.ErrGetExtra, err)
// 	}
// 	api := bone.GetValue(r, "api")
// 	req.UserId = extraInfo.UserId
// 	req.API = api
// 	return req, nil
// }

// func decodeFeedbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req feedbackRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
// 	}
// 	extraInfo, err := keto.NewExtra(r)
// 	if err != nil {
// 		return nil, errors.Wrap(keto.ErrGetExtra, err)
// 	}
// 	api := bone.GetValue(r, "api")
// 	req.UserId = extraInfo.UserId
// 	req.API = api
// 	return req, nil
// }
