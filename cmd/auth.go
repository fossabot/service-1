package cmd

import (
	"github.com/perfolio/service/internal/auth"
	"github.com/perfolio/service/internal/auth/endpoint"
	loggingMW "github.com/perfolio/service/internal/auth/logging"
	metricsMW "github.com/perfolio/service/internal/auth/metrics"
	"github.com/perfolio/service/internal/auth/model"
	"github.com/perfolio/service/internal/auth/repository"
	"github.com/perfolio/service/internal/auth/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"go.uber.org/zap"
	"github.com/spf13/cobra"
)

var port string

// authCmd represents the serve command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Serves the auth application",
	Long:  `Run this service on the specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample().With(zap.String("service", "auth"))

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			logger.Error("exit", zap.Error(err))
			os.Exit(-1)
		}
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			logger.Error("exit", zap.Error(err))
			os.Exit(-1)
		}

		svc := auth.NewService(repository.NewPostgres(db, logger), logger)
		svc = metricsMW.Use()(svc)
		svc = loggingMW.Use(logger)(svc)

		endpoints := endpoint.New(svc)
		handler := server.CreateHandler(endpoints)
		server.Run(logger, handler, port)

	},
}

func init() {
	authCmd.Flags().StringVarP(&port, "port", "p", "8000", "Run the service on this port.")
	rootCmd.AddCommand(authCmd)
}
