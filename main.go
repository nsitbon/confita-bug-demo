package main

import (
	"context"
	"fmt"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
    Port int `config:"port,required"`
	Host string `config:"host,required"`
}

func main() {
	var cfg Config
	l := confita.NewLoader(env.NewBackend(), flags.NewBackend())
	err := l.Load(context.Background(), &cfg)
	if err != nil {
		fmt.Printf("error %v\n", err)
	} else {
        fmt.Printf("ok %s:%d\n", cfg.Host, cfg.Port)
	}
}
