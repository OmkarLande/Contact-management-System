package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var store = sessions.NewCookieStore([]byte("contact-management-system"))

type AuthenticationMiddleware struct {
	ExcludedRoutes []string
}

func (mw *AuthenticationMiddleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range mw.ExcludedRoutes {
			if r.URL.Path == route {
				next.ServeHTTP(w, r)
				return
			}
		}

		userID := GetUserIDFromContextOrSession(r)

		if userID == primitive.NilObjectID {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mw *AuthenticationMiddleware) Skip(routes ...string) *AuthenticationMiddleware {
	mw.ExcludedRoutes = routes
	return mw
}

func GetUserIDFromContextOrSession(r *http.Request) primitive.ObjectID {

	session, err := store.Get(r, "contact-management-system-session")
	if err != nil {
		return primitive.NilObjectID
	}

	userID, ok := session.Values["userID"]
	if !ok {
		return primitive.NilObjectID
	}

	userIDString, ok := userID.(string)
	if !ok {
		return primitive.NilObjectID
	}

	objID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return primitive.NilObjectID
	}

	return objID
}
