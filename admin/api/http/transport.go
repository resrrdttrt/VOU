package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/viot/viot/pkg/common"
	"github.com/viot/viot/pricing"
	"net/http"
	"strings"

	"github.com/viot/viot"
	"github.com/viot/viot/mysafe"
	"github.com/viot/viot/pkg/errors"
	keto "github.com/viot/viot/pkg/messaging/ory/keto/grpc"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
)

var (
	contentType = viot.Env("CONTENT_TYPE", mysafe.ContentType)
)

var (
	errUnsupportedMediaType = errors.New("unsupported media type")
)

func MakeHandler(svc mysafe.Service) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := bone.New()

	r.Post("/mysafe/auth/:api", kithttp.NewServer(
		authEndpoint(svc),
		decodeAuthRequest,
		encodeResponse,
		opts...,
	))
	r.Get("/mysafe/token", kithttp.NewServer(
		tokenEndpoint(svc),
		decodeTokenRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/mysafe/devices/:api", kithttp.NewServer(
		deviceEndpoint(svc),
		decodeDeviceRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/mysafe/user/:api", kithttp.NewServer(
		userEndpoint(svc),
		decodeUserRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/mysafe/children/:api", kithttp.NewServer(
		childrenEndpoint(svc),
		decodeChildrenRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/mysafe/no-auth/otp", kithttp.NewServer(
		otpEndpoint(svc),
		decodeOtpRequest,
		encodeResponse,
		opts...,
	))
	r.Post("/mysafe/feedback/:api", kithttp.NewServer(
		feedbackEndpoint(svc),
		decodeFeedbackRequest,
		encodeResponse,
	))
	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", contentType)
	if ar, ok := response.(viot.Response); ok {
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
		w.Header().Set("Content-Type", contentType)
		switch {
		case errors.Contains(errorVal, mysafe.ErrNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Contains(errorVal, errUnsupportedMediaType):
			w.WriteHeader(http.StatusUnsupportedMediaType)
		case errors.Contains(errorVal, mysafe.ErrMalformedEntity):
			w.WriteHeader(http.StatusBadRequest)
		case errors.Contains(errorVal, mysafe.ErrBadRequest):
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

func decodeTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	extraInfo, err := keto.NewExtra(r)
	if err != nil {
		return nil, errors.Wrap(keto.ErrGetExtra, err)
	}
	req := tokenRequest{
		UserId: extraInfo.UserId,
	}
	return req, nil
}

func decodeAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	var req authRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
	}
	extraInfo, err := keto.NewExtra(r)
	if err != nil {
		return nil, errors.Wrap(keto.ErrGetExtra, err)
	}
	api := bone.GetValue(r, "api")
	req.UserId = extraInfo.UserId
	req.API = api
	return req, nil
}

func decodeDeviceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	var req deviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
	}
	api := bone.GetValue(r, "api")
	req.API = api
	return req, nil
}

func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
	}
	api := bone.GetValue(r, "api")
	req.API = api
	return req, nil
}

func decodeChildrenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	var req childrenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
	}
	api := bone.GetValue(r, "api")
	req.API = api
	return req, nil
}

func decodeOtpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	req := otpRequest{
		ClientKeyInfo: common.ParseClientInfo(r),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(pricing.ErrMalformedEntity, err)
	}
	return req, nil
}

func decodeFeedbackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {
		return nil, errUnsupportedMediaType
	}
	var req feedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(mysafe.ErrMalformedEntity, err)
	}
	extraInfo, err := keto.NewExtra(r)
	if err != nil {
		return nil, errors.Wrap(keto.ErrGetExtra, err)
	}
	api := bone.GetValue(r, "api")
	req.UserId = extraInfo.UserId
	req.API = api
	return req, nil
}
