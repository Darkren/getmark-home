package endpoint

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Darkren/getmark-home/pkg/data/product"
	"github.com/Darkren/getmark-home/pkg/data/user"
	"github.com/Darkren/getmark-home/pkg/service/auth"
	"github.com/Darkren/getmark-home/pkg/service/pricetag"
)

// AddProduct is the endpoint which add new product to the system.
func AddProduct(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "AddProduct"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		var p product.Product
		if err := gctx.Bind(&p); err != nil {
			gctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		p.UserID = user.ID

		if err := productsRepo.Add(&p); err != nil {
			log.Errorf("productsRepo.Add: %v\n", err)

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

// DeleteProduct is the endpoint which deletes product from the system.
func DeleteProduct(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "DeleteProduct"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		barcode := gctx.Param("barcode")

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		if err := productsRepo.Delete(barcode, user.ID); err != nil {
			log.WithFields(logrus.Fields{
				"barcode": barcode,
				"userID":  user.ID,
			}).Errorf("productsRepo.Delete: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		gctx.Status(http.StatusOK)
	}
}

// ListProducts is the endpoint which returns complete list of existing products to a client.
func ListProducts(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "ListProducts"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		products, err := productsRepo.List(user.ID)
		if err != nil {
			log.WithFields(logrus.Fields{"userID": user.ID}).Errorf("productsRepo.List: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if len(products) == 0 {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		gctx.JSON(http.StatusOK, &products)
	}
}

// GetProduct is the endpoint which returns a single product.
func GetProduct(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "GetProduct"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		barcode := gctx.Param("barcode")

		product, err := productsRepo.Product(barcode, user.ID)
		if err != nil {
			log.WithFields(logrus.Fields{
				"barcode": barcode,
				"userID":  user.ID,
			}).Errorf("productsRepo.Product: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if product == nil {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		gctx.JSON(http.StatusOK, &product)
	}
}

// GeneratePriceTag is the endpoint which generates a price tag file for the
// specified product.
func GeneratePriceTag(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository,
	priceTagService pricetag.Service) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "GeneratePriceTag"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		barcode := gctx.Param("barcode")

		product, err := productsRepo.Product(barcode, user.ID)
		if err != nil {
			log.WithFields(logrus.Fields{
				"barcode": barcode,
				"userID":  user.ID,
			}).Errorf("productsRepo.Product: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if product == nil {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		log = log.WithFields(logrus.Fields{
			"barcode": barcode,
			"name":    product.Name,
			"cost":    product.Cost,
		})

		tagContentsReader, err := priceTagService.Generate(barcode, product.Name, product.Cost)
		if err != nil {
			log.Errorf("priceTagService.Generate: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		tagContents, err := ioutil.ReadAll(tagContentsReader)
		if err != nil {
			log.Errorf("ioutil.ReadAll: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		filePath := "tags/doc_" + barcode + "_" + strconv.FormatInt(time.Now().UTC().Unix(), 10) + ".pdf"
		err = os.WriteFile(filePath, tagContents, 0644)
		if err != nil {
			log.Errorf("os.WriteFile: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		gctx.File(filePath)
	}
}

// GetPriceTag is the endpoint which returns existing generated price tag file.
func GetPriceTag(log *logrus.Logger, authService auth.Service,
	usersRepo user.Repository, productsRepo product.Repository) func(gctx *gin.Context) {
	return func(gctx *gin.Context) {
		log := log.WithFields(logrus.Fields{"endpoint": "GeneratePriceTag"})

		login, ok := mustValidateToken(gctx, log, authService)
		if !ok {
			return
		}

		log = log.WithFields(logrus.Fields{"login": login})

		user, err := usersRepo.UserByLogin(login)
		if err != nil {
			log.Errorf("usersRepo.UserByLogin: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if user == nil {
			gctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		filePath := gctx.Param("filepath")
		if filePath == "" {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		filePathTokens := strings.Split(filePath, "_")
		if len(filePathTokens) != 3 || filePathTokens[0] != "doc" {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		barcode := filePathTokens[1]

		product, err := productsRepo.Product(barcode, user.ID)
		if err != nil {
			log.WithFields(logrus.Fields{
				"barcode": barcode,
				"userID":  user.ID,
			}).Errorf("productsRepo.Product: %v\n", err)
			gctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if product == nil {
			gctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		gctx.File("tags/" + filePath)
	}
}

// mustValidateToken validates token and extracts login out of it. In case validation is failed
// request is aborted with the corresponding status.
func mustValidateToken(gctx *gin.Context, log *logrus.Entry, authService auth.Service) (string, bool) {
	tokenStr := gctx.Request.Header.Get("Authorization")
	if tokenStr == "" {
		gctx.AbortWithStatus(http.StatusBadRequest)
		return "", false
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := auth.NewToken(tokenStr)
	if err != nil {
		gctx.AbortWithStatus(http.StatusBadRequest)
		return "", false
	}

	if err := authService.ValidateToken(token); err != nil {
		gctx.AbortWithStatus(http.StatusForbidden)
		return "", false
	}

	login, err := token.Login()
	if err != nil {
		log.WithFields(logrus.Fields{"token": tokenStr}).Errorf("login is missing in the token")
		// auth service validated token, so the fact that there's no login there
		// most likely means we messed it up somewhere pretty bad
		gctx.AbortWithStatus(http.StatusInternalServerError)
		return "", false
	}

	return login, true
}
