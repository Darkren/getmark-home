package commands

import (
	"fmt"
	"github.com/Darkren/getmark-home/internal/config"
	"github.com/Darkren/getmark-home/internal/endpoint"
	"github.com/Darkren/getmark-home/pkg/api"
	"github.com/Darkren/getmark-home/pkg/data/product"
	"github.com/Darkren/getmark-home/pkg/data/user"
	"github.com/Darkren/getmark-home/pkg/service/auth"
	"github.com/Darkren/getmark-home/pkg/service/pricetag"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	listenFlag          = "listen"
	shutdownTimeoutFlag = "shutdown-timeout"
)

const (
	defaultShutdownTimeout = "5s"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the API",
	Long:  "Starts the API",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.FromEnv()
		if err != nil {
			return fmt.Errorf("config.FromEnv: %w", err)
		}

		db, err := gorm.Open(postgres.Open(cfg.DB.ToDriverDSN()))
		if err != nil {
			return fmt.Errorf("gorm.Open: %w", err)
		}

		usersRepo := user.NewPgSQLRepository(db)
		productsRepo := product.NewPgSQLRepository(db)
		priceTagService := pricetag.NewPDFService()

		tr := &http.Transport{
			MaxIdleConns:          10,
			IdleConnTimeout:       15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			DisableKeepAlives:     false,
		}
		httpCl := &http.Client{
			Transport: tr,
			// do not follow redirects
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		authService, err := auth.NewHTTPService(httpCl, cfg.AuthService.URL)
		if err != nil {
			return fmt.Errorf("auth.NewHTTPService: %w", err)
		}

		//listen := viper.GetString(listenFlag)
		shutdownTimeout := viper.GetDuration(shutdownTimeoutFlag)
		handler := endpoint.CreateRouter(logrus.New(), authService, usersRepo, productsRepo, priceTagService)

		listen := ":8081"
		return api.Run(listen, handler, api.WithShutdownTimeout(shutdownTimeout))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.InitDefaultHelpCmd()

	rootCmd.PersistentFlags().StringP(listenFlag, "l", "", "address to listen on")
	rootCmd.PersistentFlags().String(shutdownTimeoutFlag, defaultShutdownTimeout, "shutdown timeout")

	//rootCmd.MarkPersistentFlagRequired(listenFlag)

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.BindPFlags(rootCmd.Flags())
}
