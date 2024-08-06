package http

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/viot/viot/mysafe"
	"github.com/viot/viot/pkg/common"
	"github.com/viot/viot/pkg/db"
)

func tokenEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(tokenRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		if data, err := svc.Token(ctx, req.UserId); err != nil {
			return nil, err
		} else {
			return common.SuccessRes(tokenResponse{AccessToken: data}), nil
		}
	}
}

func otpEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(otpRequest)
		//if err := req.validate(); err != nil {
		//	return nil, err
		//}
		if err := svc.VhomeSendOTP(ctx, req.ClientKeyInfo, req.OTP, req.Phone); err != nil {
			return nil, err
		} else {
			return common.SuccessRes(nil), nil
		}
	}
}

func authEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		switch req.API {
		case "LinkRequest":
			if err := svc.LinkRequest(ctx, req.Phone); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(req.Phone), nil
			}
		case "LoginOTP":
			login := mysafe.VCS_LoginOTPRequest{
				Username:    req.Phone,
				Code:        req.PhoneCode,
				GrantType:   "otp",
				ClientId:    "ee6720f1b7ab3bb1a9f6",
				CountryCode: "VN",
			}
			if data, err := svc.LoginOTP(ctx, req.UserId, login); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(tokenResponse{AccessToken: data}), nil
			}
		case "Register":
			register := mysafe.VCS_RegisterRequest{
				Organization: "safe_platform_og",
				Application:  "safe_platform",
				CountryCode:  "VN",
				Phone:        req.Phone,
				Password:     generateRandomPassword(12),
				PhoneCode:    req.PhoneCode,
				AutoSignin:   true,
			}
			if data, err := svc.Register(ctx, req.UserId, register); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(tokenResponse{AccessToken: data}), nil
			}
		default:
			return nil, nil
		}
	}
}

func deviceEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deviceRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		switch req.API {
		case "GetDevice":
			queryParams := mysafe.VCS_AllDeviceRequest{
				Online: req.Online,
			}
			if data, err := svc.GetDevice(ctx, req.MacAddr, queryParams, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateDevice":
			updateDevice := mysafe.VCS_UpdateDeviceRequest{
				Name:        req.Name,
				UserAgent:   req.UserAgent,
				Os:          req.Os,
				DeviceType:  req.DeviceType,
				DeviceModel: req.DeviceModel,
			}
			if err := svc.UpdateDevice(ctx, req.MacAddr, updateDevice, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "DeleteDevice":
			if err := svc.DeleteDevice(ctx, req.MacAddr, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "AssignDeviceToChild":
			listDevice := mysafe.VCS_AssignListDeviceRequest{Devices: req.Devices}
			if err := svc.AssignDeviceToChild(ctx, req.ChildId, listDevice, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "UnassignDevice":
			if err := svc.UnassignDevice(ctx, req.MacAddr, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "InternetDevice":
			if err := svc.InternetDevice(ctx, req.MacAddr, *req.EnableInternet, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(*req.EnableInternet), nil
			}
		default:
			return nil, nil
		}
	}
}

func userEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(userRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		switch req.API {
		case "GetUserLicenses":
			if data, err := svc.GetUserLicenses(ctx, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GetFTTHContracts":
			if data, err := svc.GetFTTHContracts(ctx, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GetServicePlans":
			queryParams := mysafe.VCS_ServicePlansRequest{ServiceType: req.ServiceType}
			if data, err := svc.GetServicePlans(ctx, queryParams, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "MakeSubscription":
			subsciption := mysafe.VCS_MakeSubscriptionRequest{
				FTTH:         req.FTTH,
				ServiceType:  req.ServiceType,
				PlanId:       req.PlanId,
				RefferalCode: req.RefferalCode,
			}
			if data, err := svc.MakeSubscription(ctx, subsciption, req.Bearer); err != nil {
				fmt.Println("WTD")
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "ActivateSubscription":
			subsciption := mysafe.VCS_ActivateSubscriptionRequest{
				SupscriptionId: req.SubscriptionId,
			}
			if data, err := svc.ActivateSubscription(ctx, subsciption, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "RequestUpgradeDevice":
			requestUpgrade := mysafe.VCS_RequestUpgradeDeviceRequest{
				FTTH:       req.FTTH,
				SubId:      req.SubId,
				TicketType: req.TicketType,
			}
			if data, err := svc.RequestUpgradeDevice(ctx, requestUpgrade, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GiveFeedbackBase64":
			feedback := mysafe.VCS_GiveFeedbackBase64Request{
				Id:              req.Id,
				Rating:          req.Rating,
				Subject:         req.Subject,
				FeedbackContent: req.FeedbackContent,
				Images:          req.Images,
				CreatedAt:       req.CreatedAt,
			}
			if err := svc.GiveFeedbackBase64(ctx, feedback, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		default:
			return nil, nil
		}
	}
}

func childrenEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(childrenRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		switch req.API {
		case "GetChildren":
			queryParams := mysafe.VCS_GetAllChildrenRequest{
				Sort: req.Sort,
			}
			if data, err := svc.GetChildren(ctx, req.ChildId, queryParams, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "CreateChild":
			createChild := mysafe.VCS_CreateChildRequest{
				Name:     req.Name,
				Birthday: req.Birthday,
				Avatar:   req.Avatar,
				Gender:   req.Gender,
				Phone:    req.Phone,
				Devices:  req.Devices,
			}
			if err := svc.CreateChild(ctx, createChild, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "UpdateChild":
			updateChild := mysafe.VCS_UpdateChildRequest{
				Name:     req.Name,
				Birthday: req.Birthday,
				Avatar:   req.Avatar,
				Gender:   req.Gender,
				Phone:    req.Phone,
			}
			if err := svc.UpdateChild(ctx, updateChild, req.ChildId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "DeleteChild":
			if err := svc.DeleteChild(ctx, req.ChildId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "GetChildActivityMode":
			if data, err := svc.GetChildActivityMode(ctx, req.ChildId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildActivityMode":
			updateChildActivityMode := mysafe.VCS_UpdateChildActivityModeRequest{
				Active:         req.Active,
				FromTime:       req.FromTime,
				ToTime:         req.ToTime,
				ChildTimetable: req.ChildTimetable,
			}
			if data, err := svc.UpdateChildActivityMode(ctx, req.ChildId, updateChildActivityMode, req.ActivityModeId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GetChildAllActivities":
			getChildAllActivities := mysafe.VCS_GetChildAllActivitiesRequest{
				Start:      req.Start,
				End:        req.End,
				Limit:      req.Limit,
				Skip:       req.Skip,
				QueryBy:    req.QueryBy,
				QueryValue: req.QueryValue,
			}
			if data, err := svc.GetChildAllActivities(ctx, req.ChildId, getChildAllActivities, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GetChildOneActivity":
			if data, err := svc.GetChildOneActivity(ctx, req.ChildId, req.ActivityId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "GetChildQuickActions":
			if data, err := svc.GetChildQuickActions(ctx, req.ChildId, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildQuickActions":
			updateChildQuickActions := mysafe.VCS_UpdateChildQuickActionsRequest{
				QuickActionId:   req.QuickActionId,
				Active:          req.Active,
				DurationSeconds: req.DurationSeconds,
			}
			if data, err := svc.UpdateChildQuickActions(ctx, req.ChildId, updateChildQuickActions, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildFilterStatus":
			updateChildFilterStatus := mysafe.VCS_UpdateChildFilterStatusRequest{
				Enable: req.Enable,
			}
			if data, err := svc.UpdateChildFilterStatus(ctx, req.ChildId, req.FilterType, updateChildFilterStatus, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildAppFilter":
			updateChildAppFilter := mysafe.VCS_UpdateChildAppFilterRequest{
				AppId:  req.AppId,
				Action: req.Action,
			}
			if data, err := svc.UpdateChildAppFilter(ctx, req.ChildId, updateChildAppFilter, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildCategoryFilter":
			updateChildCategoryFilter := mysafe.VCS_UpdateChildCategoryFilterRequest{
				CategoryName: req.CategoryName,
				Action:       req.Action,
				NumId:        req.NumId,
			}
			if data, err := svc.UpdateChildCategoryFilter(ctx, req.ChildId, updateChildCategoryFilter, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "UpdateChildDomainFilter":
			updateChildDomainFilter := mysafe.VCS_UpdateChildDomainFilterRequest{
				HostName: req.HostName,
				Action:   req.Action,
			}
			if data, err := svc.UpdateChildDomainFilter(ctx, req.ChildId, updateChildDomainFilter, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "DeleteChildDomainFilter":
			deleteChildDomainFilter := mysafe.VCS_DeleteChildDomainFilterRequest{
				HostName: req.HostName,
			}
			if data, err := svc.DeleteChildDomainFilter(ctx, req.ChildId, deleteChildDomainFilter, req.Bearer); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		default:
			return nil, nil
		}
	}
}

func feedbackEndpoint(svc mysafe.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(feedbackRequest)
		if err := req.validate(); err != nil {
			return nil, err
		}
		switch req.API {
		case "GiveFeedback":
			feedback := db.MysafeUserFeedbackMongoDb{
				UserId:          req.UserId,
				Rating:          req.Rating,
				Subject:         req.Subject,
				FeedbackContent: req.FeedbackContent,
				Images:          req.Images,
			}
			if err := svc.SaveUserFeedback(ctx, feedback); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		case "GetFeedback":
			if data, err := svc.GetUserFeedback(ctx, req.UserId); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(data), nil
			}
		case "DeleteFeedback":
			if err := svc.DeleteUserFeedback(ctx, req.UserId, req.FeedbackId); err != nil {
				return nil, err
			} else {
				return common.SuccessRes(nil), nil
			}
		default:
			return nil, nil
		}
	}
}

func generateRandomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789" + "!@#$%^&*()-_=+<>?"
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
