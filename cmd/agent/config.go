package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/grnsv/metrics/internal/common"
)

type Config struct {
	ServerURL      common.NetAddress `env:"ADDRESS"`
	PollInterval   common.Time       `env:"POLL_INTERVAL"`
	ReportInterval common.Time       `env:"REPORT_INTERVAL"`
}

var config = Config{
	ServerURL:      common.NetAddress{Host: "localhost", Port: 8080},
	PollInterval:   common.Time{Duration: 2 * time.Second},
	ReportInterval: common.Time{Duration: 10 * time.Second},
}

func parseVars() error {
	flag.Var(&config.ServerURL, "a", "Address for server")
	flag.Var(&config.PollInterval, "p", "Poll interval in seconds")
	flag.Var(&config.ReportInterval, "r", "Report interval in seconds")
	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}

	return nil
}
