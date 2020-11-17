package http

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cecepsprd/crowfu-api/config"
	"github.com/cecepsprd/crowfu-api/pkg/log"
	"github.com/labstack/echo"
)

func RunServer() {
	// Load database configuration
	cfg := config.LoadConfiguration()

	// connect to database
	// db, err := config.MysqlConnect(cfg)
	// if err != nil {
	// 	log.Fatal("error connecting to database: ", err.Error())
	// }

	// init echo
	e := echo.New()

	// Starting Server
	//
	go func() {
		err := e.Start(cfg.App.HttpPort)
		if err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// block until a signal received
	sig := <-c
	log.Info("got signal :", sig)

	// gracefully shutdown the server waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	e.Shutdown(ctx)
}
