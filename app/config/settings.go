package config

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	App struct {
		Port int `env:"PORT" json:"port"`
	}
	Extras env.EnvSet `json:"-"`
}

func NewSettings() *Settings {
	var settings Settings

	extras, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}

	settings.Extras = extras
	return &settings
}

func (s Settings) BindAddress() string {
	return fmt.Sprintf(":%d", s.App.Port)
}
