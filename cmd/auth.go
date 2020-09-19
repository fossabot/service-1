package cmd

import (
	"github.com/go-kit/kit/log/level"
	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/endpoint"
	serviceMW "github.com/perfolio/service/internal/auth/middleware/service"
	"github.com/perfolio/service/internal/auth/model"
	"github.com/perfolio/service/internal/auth/repository"
	"github.com/perfolio/service/internal/auth/server"
	"github.com/perfolio/service/pkg/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"

	"github.com/spf13/cobra"
)

var port string
var iexToken string

// authCmd represents the serve command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Serves the auth application",
	Long:  `Run this service on the specified port`,
	Run: func(cmd *cobra.Command, args []string) {

		logger := logging.New("auth")

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		svc := auth.NewService(repository.NewPostgres(db, logger), logger)
		svc = serviceMW.Logging(logger)(svc)
		// svc = company.MetricsMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}

		endpoints := endpoint.New(svc, logger)
		handler := server.CreateHandler(endpoints)
		server.Run(logger, handler, port)

	},
}

func init() {
	authCmd.Flags().StringVarP(&port, "port", "p", "8000", "Run the service on this port.")
	rootCmd.AddCommand(authCmd)
}
