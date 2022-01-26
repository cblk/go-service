package cmds

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cblk/go-service/api"
	"github.com/cblk/go-service/config"
	"github.com/cblk/go-service/internal/service/db"
	"github.com/cblk/go-service/internal/service/origin"
	"github.com/cblk/go-service/internal/service/session"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	RunE: func(cmd *cobra.Command, args []string) error {
		InitServerFromAppConfig()
		conf := config.GetConfig()
		logrus.Info("start service server")
		app := api.GetHttpApplication(conf)
		address := fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port)
		logrus.Info("server url:" + address)

		return app.Run(address)
	},
}

func InitServerFromAppConfig() {
	appConfig := config.GetConfig()

	// Initialize database
	if err := db.InitDB(appConfig); err != nil {
		panic(err)
	}
	// Initialize session store
	if err := session.InitSessionStore(appConfig); err != nil {
		panic(err)
	}
	// Initialize cors allow origin
	if err := origin.InitOrigin(appConfig); err != nil {
		panic(err)
	}
}

func gracefulShutDownServer(address string, engine http.Handler) {
	srv := &http.Server{
		Addr:    address,
		Handler: engine,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown: ", err)
	}
	logrus.Info("Server exit")
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}
