package middlewares

import (
	"net/http"

	"github.com/resrrdttrt/VOU/admin"
)

// func VerifyAdminMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		accessToken := r.Header.Get("Authorization")
// 		if accessToken == "" {
// 			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
// 			return
// 		}
// 		userID, err := admin.GetUserIDByAccessToken(accessToken)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusUnauthorized)
// 			return
// 		}
// 		role, err := admin.GetUserRoleByID(userID)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusUnauthorized)
// 			return
// 		}
// 		if role != "admin" {
// 			http.Error(w, "You are not authorized to access this resource", http.StatusForbidden)
// 			return
// 		}
// 		next.ServeHTTP(w, r)

// 	})
// }

func VerifyEventMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		userID, err := admin.GetUserIDByAccessToken(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		role, err := admin.GetUserRoleByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if role != "enterprise" {
			http.Error(w, "You are not authorized to access this resource", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}
