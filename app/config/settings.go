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
	Database struct {
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Port     int    `env:"DATABASE_PORT"`
		Name     string `env:"DATABASE_NAME"`
	}
	External struct {
		SearchEndpoint string `env:"NAVER_SEARCH_API"`
		ClientID       string `env:"NAVER_CLIENT_ID"`
		ClientSecret   string `env:"NAVER_CLIENT_SECRET"`
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
