package main

import "github.com/joeshaw/envdecode"

type Config struct {
    Port        string  `env:"PORT,default=80"`
    LineSecret  string  `env:"LINE_SECRET,required"`
    LineToken   string  `env:"LINE_TOKEN,required"`
}

func (c *Config) load() error {
    return envdecode.Decode(c)
}
