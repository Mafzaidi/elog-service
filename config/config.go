package config

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server *Server
		App    *App
		DB     *DB
		JWT    *JWT
	}

	App struct {
		Name    string
		Version string
	}

	Server struct {
		Host string
		Port int
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}

	JWT struct {
		Secret        string
		TokenExpiry   time.Duration
		RefreshExpiry time.Duration
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic("Failed to read config.yaml: " + err.Error())
		}

		cfg := &Config{
			Server: &Server{},
			App:    &App{},
			DB:     &DB{},
			JWT:    &JWT{},
		}

		if err := viper.Unmarshal(cfg); err != nil {
			panic("Failed to unmarshal config into struct: " + err.Error())
		}

		_ = godotenv.Load()

		cfg.DB.Host = getEnvOrDefault("MONGO_DB_HOST", cfg.DB.Host)
		cfg.DB.Port = getEnvOrDefault("MONGO_DB_PORT", cfg.DB.Port)
		cfg.DB.User = getEnvOrDefault("MONGO_USER", cfg.DB.User)
		cfg.DB.Password = getEnvOrDefault("MONGO_PASSWORD", cfg.DB.Password)
		cfg.DB.DBName = getEnvOrDefault("MONGO_DB_NAME", cfg.DB.DBName)

		cfg.JWT.Secret = getEnvOrDefault("JWT_SECRET", cfg.JWT.Secret)

		if s := viper.GetString("jwt.tokenExpiry"); s != "" {
			cfg.JWT.TokenExpiry, _ = time.ParseDuration(s)
		}
		if s := viper.GetString("jwt.refreshExpiry"); s != "" {
			cfg.JWT.RefreshExpiry, _ = time.ParseDuration(s)
		}

		configInstance = cfg
	})

	return configInstance
}

func getEnvOrDefault(envKey, fallback string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return fallback
}
