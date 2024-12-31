package main

import (
	"flag"

	"github.com/grnsv/metrics/internal/common"
)

var (
	address = common.NetAddress{Host: "localhost", Port: 8080}
)

func ParseFlags() {
	flag.Var(&address, "a", "Address for server")
	flag.Parse()
}
