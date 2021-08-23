package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/controller"
	_ "github.com/betorvs/playbypost-dnd/gateway/customlog"
	_ "github.com/betorvs/playbypost-dnd/gateway/mongodb"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	// from: https://echo.labstack.com/cookbook/graceful-shutdown/
	// Start server
	go func() {
		// change ready probe to UP
		config.Values.IsReady.Store(true)
		if err := e.Start(":" + config.Values.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logLocal := config.GetLogger()
	signal := <-quit
	logLocal.Info("Received signal ", signal)
	config.Values.IsReady.Store(false)
	if config.Values.LogLevel == "DEBUG" {
		time.Sleep(2 * time.Second)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Values.WaitTime)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
