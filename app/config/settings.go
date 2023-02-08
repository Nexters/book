package config

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

// Settings 환경변수
type Settings struct {
	// App 앱 환경변수
	App struct {
		Port     int    `env:"PORT" json:"port"`
		API_HOST string `env:"API_HOST"`
	}

	// Database DB 환경변수
	Database struct {
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Port     int    `env:"DATABASE_PORT"`
		Name     string `env:"DATABASE_NAME"`
		URL      string `env:"DATABASE_URL"`
	}

	// External 외부 연동 환경변수
	External struct {
		SearchEndpoint string `env:"NAVER_SEARCH_API"`
		ClientID       string `env:"NAVER_CLIENT_ID"`
		ClientSecret   string `env:"NAVER_CLIENT_SECRET"`
	}
	Extras env.EnvSet `json:"-"`
}

// NewSettings 생성자
func NewSettings() *Settings {
	var settings Settings

	extras, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}

	settings.Extras = extras
	return &settings
}

// BindAddress 포트 연동
func (s Settings) BindAddress() string {
	return fmt.Sprintf(":%d", s.App.Port)
}
