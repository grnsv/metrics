package common

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type NetAddress struct {
	Host string
	Port int
}

func (a NetAddress) String() string {
	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("address must be in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}

func (a *NetAddress) UnmarshalText(text []byte) error {
	return a.Set(string(text))
}

type Time struct {
	Duration time.Duration
}

func (t Time) String() string {
	return t.Duration.String()
}

func (t *Time) Set(s string) error {
	duration, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	t.Duration = time.Duration(duration) * time.Second
	return nil
}

func (t *Time) UnmarshalText(text []byte) error {
	return t.Set(string(text))
}
