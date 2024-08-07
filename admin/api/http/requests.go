package http

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/resrrdttrt/VOU/pkg/errors"
)

var (
	ErrInvalidUUID      = errors.New("invalid uuid")
	ErrInvalidRoleValue = errors.New("role must be enterprise, end_user or admin")
	ErrInvalidStatus    = errors.New("status must be active or inactive")
)

func errMissing(field string) error {
	return errors.Wrap(errors.ErrMalformedEntity, errors.New(fmt.Sprintf("missing field `%s`", field)))
}

type createUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func (req createUserRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	if req.Username == "" {
		return errMissing("username")
	}
	if req.Password == "" {
		return errMissing("password")
	}
	if req.Email == "" {
		return errMissing("email")
	}
	if req.Phone == "" {
		return errMissing("phone")
	}
	if req.Role == "" {
		return errMissing("role")
	}
	if req.Role != "enterprise" && req.Role != "end_user" && req.Role != "admin" {
		return ErrInvalidRoleValue
	}
	if req.Status == "" {
		return errMissing("status")
	}
	if req.Status != "active" && req.Status != "inactive" {
		return ErrInvalidStatus
	}
	return nil
}

type updateUserRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (req updateUserRequest) validate() error {
	if req.ID == "" {
		return errMissing("user_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	if req.Role != "" && req.Role != "enterprise" && req.Role != "end_user" && req.Role != "admin" {
		return ErrInvalidRoleValue
	}
	if req.Status != "" && req.Status != "active" && req.Status != "inactive" {
		return ErrInvalidStatus
	}
	return nil
}

type getUserRequest struct {
	ID string `json:"id"`
}

func (req getUserRequest) validate() error {
	if req.ID == "" {
		return errMissing("user_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}


type getGameRequest struct {
	ID string `json:"id"`
}

func (req getGameRequest) validate() error {
	if req.ID == "" {
		return errMissing("game_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type createGameRequest struct {
	Name        string `json:"name"`
	Images	  string `json:"images"`
	Type 	  string `json:"type"`
	ExchangeAllow bool `json:"exchange_allow"`
	Tutorial  string `json:"tutorial"`
}

func (req createGameRequest) validate() error {
	if req.Name == "" {
		return errMissing("name")
	}
	if req.Images == "" {
		return errMissing("images")
	}
	if req.Type == "" {
		return errMissing("type")
	}
	if req.Tutorial == "" {
		return errMissing("tutorial")
	}
	return nil
}

type updateGameRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Images	  string `json:"images"`
	Type 	  string `json:"type"`
	ExchangeAllow bool `json:"exchange_allow"`
	Tutorial  string `json:"tutorial"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (req updateGameRequest) validate() error {
	if req.ID == "" {
		return errMissing("game_id")
	} else {
		if _, err := uuid.Parse(req.ID); err != nil {
			return errors.Wrap(errors.ErrMalformedEntity, ErrInvalidUUID)
		}
	}
	return nil
}

type statisticInTimeRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (req statisticInTimeRequest) validate() error {
	if req.Start.IsZero() {
		return errMissing("start")
	}
	if req.End.IsZero() {
		return errMissing("end")
	}
	return nil
}