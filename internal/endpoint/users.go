package endpoint

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Darkren/getmark-home/pkg/data/user"
	"github.com/Darkren/getmark-home/pkg/service/auth"
)

// AddUser is the endpoint registering new user in the system.
func AddUser(log *logrus.Logger, usersRepo user.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "AddUser"})

		var u user.User
		if err := gctx.Bind(&u); err != nil {
			gctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := usersRepo.Add(&u); err != nil {
			log.Errorf("userRepo.Add: %v\n", err)

			if strings.Contains(err.Error(), "duplicate key") {
				gctx.AbortWithStatus(http.StatusConflict)
				return
			}

			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		gctx.Status(http.StatusOK)
	}
}

// AuthRequest is the authentication request.
type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthResponse is the authentication response.
type AuthResponse struct {
	Token string `json:"token"`
}

// Auth is the endpoint performing authentication of a user in the system.
func Auth(log *logrus.Logger, authService auth.Service, usersRepo user.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "Auth"})

		var req AuthRequest
		if err := gctx.Bind(&req); err != nil {
			gctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		log = log.WithFields(logrus.Fields{"login": req.Login})

		user, err := usersRepo.UserByLogin(req.Login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if user == nil {
			gctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Password != req.Password {
			gctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := authService.Auth(req.Login)
		if err != nil {
			log.Errorf("authService.Auth: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		resp := AuthResponse{
			Token: token.Token,
		}

		gctx.JSON(http.StatusOK, &resp)
	}
}
