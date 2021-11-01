package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/log"
	"github.com/arimaulana/point-of-no-return/internal/sample/configs"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New()
	if len(os.Args) <= 2 {
		logger.Error("Usage: ", os.Args[1], " command", " argument")
		return errors.New("invalid command")
	}

	cfg, err := configs.LoadConfig()
	if err != nil {
		logger.Errorf("failed to load application configuration")
		return err
	}

	switch os.Args[1] {
	case "migrate":
		err = MigrateSql(cfg, logger, os.Args[2])
	case "seed":
		err = SeedSql(cfg, logger, os.Args[2])
	default:
		err = errors.New("must specify a command")
	}

	if err != nil {
		return err
	}

	return nil
}
