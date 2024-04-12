package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Logger struct {
	LoggerFilePath  string
	LoggerFileFlag  bool
	LoggerMultiFlag bool
}

type NetAddress struct {
	Host string
	Port int
}

type Storage struct {
	DatabaseDSN string
}

type Caches struct {
	Url string
}

type Flags struct {
	Logger
	NetAddress
	Storage
	Caches
}

func (a NetAddress) String() string {
	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return errors.New("need address in a form host:port")
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return fmt.Errorf("cannot atoi port: %w", err)
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}
