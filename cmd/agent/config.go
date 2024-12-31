package main

import (
	"flag"
	"time"

	"github.com/grnsv/metrics/internal/common"
)

var (
	serverURL      = common.NetAddress{Host: "localhost", Port: 8080}
	pollInterval   = common.Time{Duration: 2 * time.Second}
	reportInterval = common.Time{Duration: 10 * time.Second}
)

func ParseFlags() {
	flag.Var(&serverURL, "a", "Address for server")
	flag.Var(&pollInterval, "p", "Poll interval in seconds")
	flag.Var(&reportInterval, "r", "Report interval in seconds")
	flag.Parse()
}
