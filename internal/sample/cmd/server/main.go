package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/arimaulana/point-of-no-return/internal/sample/cmd/server/http"
	"github.com/arimaulana/point-of-no-return/internal/sample/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

func main() {
	flag.Parse()

	// create root logger tagged with server version
	logger := log.New().With(context.Background(), "version", Version)

	cfg, err := configs.LoadConfig()
	if err != nil {
		logger.Errorf("cannot load config: %s", err)
		os.Exit(1)
	}

	dsn := generateDSN(cfg)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Errorf("failed to connect to database: %s", err)
		os.Exit(1)
	}

	serverType := strings.ToLower(cfg.ServerToRun)
	switch serverType {
	case "http":
		httpServer, err := http.NewHttpServer(db, logger)
		if err != nil {
			logger.Errorf("cannot prepare server: %s", err)
			os.Exit(1)
		}

		httpServer.Run(cfg, logger)
	case "grpc":
		logger.Error("server type 'grpc' is not implemented in this service")
		os.Exit(1)
	default:
		logger.Errorf("server type '%s' is not supported", serverType)
		os.Exit(1)
	}
}

func generateDSN(cfg configs.Config) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
}
