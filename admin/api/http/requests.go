package http

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/resrrdttrt/VOU/pkg/errors"
)

var (
	ErrInvalidUUID      = errors.New("invalid uuid")
	ErrInvalidRoleValue = errors.New("role must be enterprise, end_user or admin")
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
