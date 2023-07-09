package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Darkren/getmark-home/pkg/data/user"
	"github.com/Darkren/getmark-home/pkg/service/auth"
)

// CreateRouter creates router with all the API endpoints.
func CreateRouter(log *logrus.Logger, authService auth.Service, usersRepo user.Repository) http.Handler {
	router := gin.Default()

	router.POST("/users/", AddUser(log, usersRepo))
	router.POST("/auth", Auth(log, authService, usersRepo))

	return router
}
