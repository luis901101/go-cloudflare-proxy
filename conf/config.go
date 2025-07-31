package conf

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Account     string  `env:"ACCOUNT,required"`
	Token       string  `env:"TOKEN,required"`
	ApiKey      string  `env:"API_KEY,required"`
	WatermarkId *string `env:"WATERMARK_ID"`
}

func (config Config) BaseUrl() string {
	return "https://api.cloudflare.com/client/v4/accounts/" + config.Account
}

func (config Config) Authorization() string {
	return "Bearer " + config.Token
}

func (config Config) ImageUrl() string {
	return config.BaseUrl() + "/images/v1"
}

func (config Config) StreamUrl() string {
	return config.BaseUrl() + "/stream"
}

func (config Config) DirectUploadImageUrl() string {
	return config.BaseUrl() + "/images/v2/direct_upload"
}

func (config Config) DirectUploadStreamUrl() string {
	return config.BaseUrl() + "/stream/direct_upload"
}

func (config Config) DirectUploadTusStreamUrl() string {
	return config.BaseUrl() + "/stream?direct_user=true"
}

func (config Config) WebHookUrl() string {
	return config.BaseUrl() + "/stream/webhook"
}

func Load() Config {
	// 1. Load .env file for local development
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %e", err)
	}

	// 2. Parse environment variables into the config struct
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse configuration: %+e", err)
	}

	// 3. Return the loaded configuration
	return cfg
}
