package http

import (
	"fmt"
	"github.com/viot/viot/pkg/common"
	"net"
	"time"

	"github.com/viot/viot/mysafe"
	"github.com/viot/viot/pkg/errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidMacAddr            = errors.New("invalid mac address")
	ErrInvalidUUID               = errors.New("invalid uuid")
	ErrInvalidAPI                = errors.New("invalid api")
	ErrServiceTypeValue          = errors.New("service_type must be HOME or NET")
	ErrFilterTypeValue           = errors.New("filter_type must be CATEGORIES, DOMAINS or APPS")
	ErrAppFilterActionValue      = errors.New("app filter action must be block or allow")
	ErrDomainFilterActionValue   = errors.New("domain filter action must be block or alert")
	ErrCategoryFilterActionValue = errors.New("category filter action must be block or allow")
)

func errMissing(field string) error {
	return errors.Wrap(mysafe.ErrMalformedEntity, errors.New(fmt.Sprintf("missing field `%s`", field)))
}

type otpRequest struct {
	ClientKeyInfo common.ClientKeyInfo
	OTP           string `json:"otp"`
	Phone         string `json:"phone"`
}

func (req otpRequest) validate() error {
	if req.ClientKeyInfo.ClientKey == "" {
		return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("ClientKey"))
	}
	if req.ClientKeyInfo.Signature == "" {
		return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("Signature"))
	}
	if req.Phone == "" {
		return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("phone"))
	}
	if req.OTP == "" {
		return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("otp"))
	}
	return nil
}

type tokenRequest struct {
	UserId string `json:"user_id"`
}

func (req tokenRequest) validate() error {
	if req.UserId == "" {
		return mysafe.ErrUnauthorized
	}
	return nil
}

type authRequest struct {
	API           string
	UserId        string
	Organization  string `json:"organization"`
	Application   string `json:"application"`
	CountryCode   string `json:"countryCode"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	PhoneCode     string `json:"phoneCode"`
	Code          string `json:"code"`
	GrantType     string `json:"grant_type"`
	Username      string `json:"username"`
	RefreshToken  string `json:"refresh_token"`
	AutoSignin    bool   `json:"auto_signin"`
	Type          string `json:"type"`
	ApplicationId string `json:"applicationId"`
	Method        string `json:"method"`
	CaptchaType   string `json:"captchaType"`
}

func (req authRequest) validate() error {
	if req.UserId == "" {
		return mysafe.ErrUnauthorized
	}
	if req.Phone == "" {
		return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("phone"))
	}
	switch req.API {
	case "LinkRequest":
		return nil
	case "Register":
		if req.PhoneCode == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("phoneCode"))
		}
	case "LoginOTP":
		if req.PhoneCode == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("phoneCode"))
		}
	default:
		return errors.Wrap(mysafe.ErrBadRequest, ErrInvalidAPI)
	}
	return nil
}

type deviceRequest struct {
	API            string
	MacAddr        string   `json:"mac_addr"`
	Name           string   `json:"name"`
	UserAgent      string   `json:"user_agent"`
	Os             string   `json:"os"`
	DeviceType     string   `json:"device_type"`
	DeviceModel    string   `json:"device_model"`
	Devices        []string `json:"devices"`
	ChildId        string   `json:"child_id"`
	EnableInternet *bool    `json:"enable_internet"`
	Online         bool     `json:"online"`
	Bearer         string   `json:"access_token"`
}

func (req deviceRequest) validate() error {
	if req.Bearer == "" {
		return mysafe.ErrUnauthorized
	}
	switch req.API {
	case "GetDevice":
		if req.MacAddr != "" {
			if _, err := net.ParseMAC(req.MacAddr); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
	case "UpdateDevice":
		if req.MacAddr == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("mac_addr"))
		} else {
			if _, err := net.ParseMAC(req.MacAddr); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
	case "DeleteDevice":
		if req.MacAddr == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("mac_addr"))
		} else {
			if _, err := net.ParseMAC(req.MacAddr); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
	case "AssignDeviceToChild":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.Devices == nil {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("devices"))
		} else {
			for _, mac := range req.Devices {
				if _, err := net.ParseMAC(mac); err != nil {
					return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
				}
			}
		}
	case "UnassignDevice":
		if req.MacAddr == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("mac_addr"))
		} else {
			if _, err := net.ParseMAC(req.MacAddr); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
	case "InternetDevice":
		if req.MacAddr == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("mac_addr"))
		} else {
			if _, err := net.ParseMAC(req.MacAddr); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
		if req.EnableInternet == nil {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("enable_internet"))
		}
	default:
		return errors.Wrap(mysafe.ErrBadRequest, ErrInvalidAPI)
	}
	return nil
}

type userRequest struct {
	API             string
	FTTH            string    `json:"ftth"`
	ServiceType     string    `json:"service_type"`
	PlanId          string    `json:"plan_id"`
	RefferalCode    string    `json:"refferal_code"`
	SubscriptionId  string    `json:"subscription_id"`
	SubId           string    `json:"sub_id"`
	TicketType      string    `json:"ticket_type"`
	Id              string    `json:"id"`
	Rating          int       `json:"rating"`
	Subject         string    `json:"subject"`
	FeedbackContent string    `json:"feedback_content"`
	Images          []string  `json:"images"`
	CreatedAt       time.Time `json:"created_at"`
	Bearer          string    `json:"access_token"`
}

func (req userRequest) validate() error {
	if req.Bearer == "" {
		return mysafe.ErrUnauthorized
	}
	switch req.API {
	case "GetUserLicenses":
		return nil
	case "GetFTTHContracts":
		return nil
	case "GetServicePlans":
		if req.ServiceType == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("service_type"))
		} else if req.ServiceType != "HOME" && req.ServiceType != "NET" {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrServiceTypeValue)
		}
	case "MakeSubscription":
		if req.PlanId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("plan_id"))
		}
		if req.ServiceType == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("service_type"))
		}
		if req.FTTH == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("ftth"))
		}
	case "ActivateSubscription":
		if req.SubscriptionId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("subscription_id"))
		}
	case "RequestUpgradeDevice":
		if req.SubId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("sub_id"))
		}
		if req.FTTH == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("ftth"))
		}
	case "GiveFeedbackBase64":
		if req.Subject == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("subject"))
		}
		if req.Rating < 1 || req.Rating > 5 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("rating"))
		}
		if req.CreatedAt == (time.Time{}) {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("created_at"))
		}
	default:
		return errors.Wrap(mysafe.ErrBadRequest, ErrInvalidAPI)
	}
	return nil
}

type childrenRequest struct {
	API             string
	ChildId         string                  `json:"child_id"`
	Name            string                  `json:"name"`
	Birthday        time.Time               `json:"birthday"`
	Gender          string                  `json:"gender"`
	Avatar          string                  `json:"avatar"`
	Phone           string                  `json:"phone"`
	Devices         []mysafe.Device         `json:"devices"`
	Active          bool                    `json:"active"`
	FromTime        string                  `json:"from_time"`
	ToTime          string                  `json:"to_time"`
	ChildTimetable  []mysafe.ChildTimetable `json:"child_timetable"`
	QuickActionId   string                  `json:"quick_action_id"`
	DurationSeconds int                     `json:"duration_seconds"`
	ActivityModeId  string                  `json:"activity_mode_id"`
	Enable          bool                    `json:"enable"`
	FilterType      string                  `json:"filter_type"`
	AppId           string                  `json:"app_id"`
	Action          string                  `json:"action"`
	CategoryName    string                  `json:"category_name"`
	NumId           int                     `json:"num_id"`
	HostName        string                  `json:"host_name"`
	Sort            string                  `json:"sort"`
	Start           int64                   `json:"start"`
	End             int64                   `json:"end"`
	Limit           int                     `json:"limit"`
	Skip            int                     `json:"skip"`
	QueryBy         string                  `json:"query_by"`
	QueryValue      string                  `json:"query_value"`
	ActivityId      string                  `json:"activity_id"`
	Bearer          string                  `json:"access_token"`
}

func (req childrenRequest) validate() error {
	if req.Bearer == "" {
		return mysafe.ErrUnauthorized
	}
	switch req.API {
	case "GetChildren":
		if req.ChildId != "" {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "CreateChild":
		for _, device := range req.Devices {
			if _, err := net.ParseMAC(device.Mac); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidMacAddr)
			}
		}
	case "UpdateChild":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "DeleteChild":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "GetChildActivityMode":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "UpdateChildActivityMode":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "GetChildQuickActions":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "UpdateChildQuickActions":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.QuickActionId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("quick_action_id"))
		} else {
			if _, err := uuid.Parse(req.QuickActionId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
	case "GetChildAllActivities":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.Start == 0 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("start"))
		}
		if req.End == 0 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("end"))
		}
		if req.Limit == 0 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("limit"))
		}
		if req.Skip == 0 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("skip"))
		}
	case "GetChildOneActivity":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.ActivityId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("activity_id"))
		}
	case "UpdateChildFilterStatus":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.FilterType == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("filter_type"))
		} else if req.FilterType != "CATEGORIES" && req.FilterType != "DOMAINS" && req.FilterType != "APPS" {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrFilterTypeValue)
		}
		// TODO: enable
	case "UpdateChildAppFilter":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.AppId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("app_id"))
		} else {
			if _, err := uuid.Parse(req.AppId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.Action == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("action"))
		} else if req.Action != "block" && req.Action != "allow" {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrAppFilterActionValue)
		}
	case "UpdateChildDomainFilter":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.HostName == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("host_name"))
		}
		if req.Action == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("action"))
		} else if req.Action != "block" && req.Action != "alert" {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrDomainFilterActionValue)
		}
	case "UpdateChildCategoryFilter":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.CategoryName == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("category_name"))
		}
		if req.Action == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("action"))
		} else if req.Action != "block" && req.Action != "allow" {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrCategoryFilterActionValue)
		}
		if req.NumId == 0 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("num_id"))
		}
	case "DeleteChildDomainFilter":
		if req.ChildId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("child_id"))
		} else {
			if _, err := uuid.Parse(req.ChildId); err != nil {
				return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
			}
		}
		if req.HostName == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("host_name"))
		}
	default:
		return errors.Wrap(mysafe.ErrBadRequest, ErrInvalidAPI)
	}
	return nil
}

type feedbackRequest struct {
	API             string
	Rating          int      `json:"rating"`
	FeedbackContent string   `json:"feedback_content"`
	Subject         string   `json:"subject"`
	Images          []string `json:"images"`
	UserId          string   `json:"user_id"`
	FeedbackId      string   `json:"feedback_id"`
}

func (req feedbackRequest) validate() error {
	switch req.API {
	case "GiveFeedback":
		if req.Rating < 1 || req.Rating > 5 {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("rating"))
		}
	case "GetFeedback":
		return nil
	case "DeleteFeedback":
		if req.FeedbackId == "" {
			return errors.Wrap(mysafe.ErrMalformedEntity, errMissing("feedback_id"))
		}
		if _, err := uuid.Parse(req.FeedbackId); err != nil {
			return errors.Wrap(mysafe.ErrMalformedEntity, ErrInvalidUUID)
		}
	default:
		return errors.Wrap(mysafe.ErrBadRequest, ErrInvalidAPI)
	}
	return nil
}
