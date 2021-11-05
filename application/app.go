package application

import (
	"context"
	"github.com/foxfurry/go_delivery/internal/http/controller"
	"github.com/foxfurry/go_delivery/internal/infrastracture/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type deliveryApp struct {
	server *http.Server
}

func Create(ctx context.Context) IApp {
	appHandler := gin.New()

	ctrl := controller.NewDeliveryController()
	ctrl.RegisterDeliveryRoutes(appHandler)

	app := deliveryApp{
		server: &http.Server{
			Addr:              viper.GetString("delivery_host"),
			Handler:           appHandler,
		},
	}

	return &app
}

func (d *deliveryApp) Start() {
	logger.LogMessage("Starting delivery server")

	if err := d.server.ListenAndServe(); err != http.ErrServerClosed {
		logger.LogPanicF("Unexpected error while running server: %v", err)
	}
}

func (d *deliveryApp) Shutdown() {
	if err := d.server.Shutdown(context.Background()); err != nil {
		logger.LogPanicF("Unexpected error while closing server: %v", err)
	}
	logger.LogMessage("Server terminated successfully")
}
