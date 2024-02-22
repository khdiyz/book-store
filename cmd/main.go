package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/khdiyz/book-store/config"
	_ "github.com/khdiyz/book-store/docs"
	"github.com/khdiyz/book-store/pkg/handler"
	"github.com/khdiyz/book-store/pkg/repository"
	"github.com/khdiyz/book-store/pkg/service"
	"github.com/khdiyz/book-store/server"
	"github.com/khdiyz/book-store/utils/logger"
)

// @title Book-Store System
// @version 1.0
// @description API Server for Book-Store Application
// @host localhost:9000
// @BasePath
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := logger.GetLogger()
	cfg := config.Load()

	err := repository.MigrateSQLite(cfg, log, true)
	if err != nil {
		log.Error("error while migrate up", err)
	}

	db, err := repository.NewSQLiteDB(repository.Config{
		FilePath: cfg.SQLiteFilePath,
	})
	if err != nil {
		log.Fatal("failed to initialize db:", err)
		return
	}

	log.Info("Successfully connected to database")

	repos := repository.NewRepository(db, log)
	services := service.NewService(repos, log)
	handlers := handler.NewHandler(services, log)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.ServerPort, handlers.InitRoutes()); err != nil {
			log.Fatal("error occured while running http server:", err)
		}
	}()

	log.Info("Book-Store App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("error occured on server shutting down:", err)
	}

	db.Close()
}
