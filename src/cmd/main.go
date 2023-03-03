package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/daniiarov-alym/micro-template/src/app/api"
	"github.com/daniiarov-alym/micro-template/src/app/repository"
	"github.com/daniiarov-alym/micro-template/src/app/service"
	"github.com/daniiarov-alym/micro-template/src/config"

	"syscall"

	"github.com/daniiarov-alym/micro-template/src/db"

	logger "github.com/sirupsen/logrus"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.TraceLevel)
	logger.SetReportCaller(true)
	config.ReadConfig()
	cfg := config.Config()
	pgxpool := db.StartDatabase(ctx)
	cmdRepo := repository.NewServerRepository(pgxpool)
	cmdService := service.NewCommandService(cmdRepo)
	server := apiserver.NewServer(cmdService)
	go func() {
		err := server.Start(cfg.ApiPort)
		if err != nil {
			logger.Fatal("failed to start due to: " + err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("Got shutdown signal")
		cancel()
	}()

	logger.Info("Initialized command service")
	<-ctx.Done()
}
