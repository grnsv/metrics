package main

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/grnsv/metrics/internal/common"
)

type Config struct {
	Address common.NetAddress `env:"ADDRESS"`
}

var config = Config{
	Address: common.NetAddress{Host: "localhost", Port: 8080},
}

func parseVars() error {
	flag.Var(&config.Address, "a", "Address for server")
	flag.Parse()

	err := env.Parse(&config)
	if err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}

	return nil
}
