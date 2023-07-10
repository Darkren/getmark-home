package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Darkren/getmark-home/pkg/data"
	"github.com/Darkren/getmark-home/pkg/data/schema"
)

func TestAddUser(t *testing.T) {
	t.Run("400 for invalid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		gctx, _ := gin.CreateTestContext(w)

		gctx.Request = httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader([]byte("invalid")))
		gctx.Request.Header.Add("Content-Type", "application/json")

		AddUser(logrus.New(), nil)(gctx)

		assert.Equal(t, http.StatusBadRequest, gctx.Writer.Status())
	})

	t.Run("test 400 for missing fields in request", func(t *testing.T) {
		w := httptest.NewRecorder()
		gctx, _ := gin.CreateTestContext(w)

		reqBody := AddUserRequest{
			Login:    "test login",
			Password: "test password",
			Email:    "test@test.test",
		}
		reqBodyJSON, err := json.Marshal(&reqBody)
		require.NoError(t, err)

		gctx.Request = httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(reqBodyJSON))
		gctx.Request.Header.Add("Content-Type", "application/json")

		AddUser(logrus.New(), nil)(gctx)

		assert.Equal(t, http.StatusBadRequest, gctx.Writer.Status())
	})

	t.Run("test 409 for already existing user", func(t *testing.T) {
		w := httptest.NewRecorder()
		gctx, _ := gin.CreateTestContext(w)

		reqBody := AddUserRequest{
			Login:    "test login",
			Password: "test password",
			Name:     "test name",
			Email:    "test@test.test",
		}
		reqBodyJSON, err := json.Marshal(&reqBody)
		require.NoError(t, err)

		gctx.Request = httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(reqBodyJSON))
		gctx.Request.Header.Add("Content-Type", "application/json")

		usersRepo := &data.MockUserRepository{}
		usersRepo.On("Add", &schema.User{
			Login:    reqBody.Login,
			Password: reqBody.Password,
			Name:     reqBody.Name,
			Email:    reqBody.Email,
		}).Return(fmt.Errorf("duplicate key"))

		AddUser(logrus.New(), usersRepo)(gctx)

		assert.Equal(t, http.StatusConflict, gctx.Writer.Status())
	})

	t.Run("test 500 on unexpected data layer error", func(t *testing.T) {
		w := httptest.NewRecorder()
		gctx, _ := gin.CreateTestContext(w)

		reqBody := AddUserRequest{
			Login:    "test login",
			Password: "test password",
			Name:     "test name",
			Email:    "test@test.test",
		}
		reqBodyJSON, err := json.Marshal(&reqBody)
		require.NoError(t, err)

		gctx.Request = httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(reqBodyJSON))
		gctx.Request.Header.Add("Content-Type", "application/json")

		usersRepo := &data.MockUserRepository{}
		usersRepo.On("Add", &schema.User{
			Login:    reqBody.Login,
			Password: reqBody.Password,
			Name:     reqBody.Name,
			Email:    reqBody.Email,
		}).Return(fmt.Errorf("unexpected error"))

		AddUser(logrus.New(), usersRepo)(gctx)

		assert.Equal(t, http.StatusInternalServerError, gctx.Writer.Status())
	})

	t.Run("test 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		gctx, _ := gin.CreateTestContext(w)

		reqBody := AddUserRequest{
			Login:    "test login",
			Password: "test password",
			Name:     "test name",
			Email:    "test@test.test",
		}
		reqBodyJSON, err := json.Marshal(&reqBody)
		require.NoError(t, err)

		gctx.Request = httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(reqBodyJSON))
		gctx.Request.Header.Add("Content-Type", "application/json")

		usersRepo := &data.MockUserRepository{}
		usersRepo.On("Add", &schema.User{
			Login:    reqBody.Login,
			Password: reqBody.Password,
			Name:     reqBody.Name,
			Email:    reqBody.Email,
		}).Return(error(nil))

		AddUser(logrus.New(), usersRepo)(gctx)

		assert.Equal(t, http.StatusOK, gctx.Writer.Status())
	})
}
