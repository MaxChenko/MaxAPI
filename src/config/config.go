package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:",squash"`
	Postgres PostgresConfig `mapstructure:",squash"`
	Redis    RedisConfig    `mapstructure:",squash"`
	Password PasswordConfig `mapstructure:",squash"`
	Cors     CorsConfig     `mapstructure:",squash"`
	Logger   LoggerConfig   `mapstructure:",squash"`
	Otp      OtpConfig      `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
}

type ServerConfig struct {
	InternalPort string `mapstructure:"server_internal_port"`
	ExternalPort string `mapstructure:"server_external_port"`
	RunMode      string `mapstructure:"server_run_mode"`
	Domain       string `mapstructure:"server_domain"`
}

type LoggerConfig struct {
	FilePath string `mapstructure:"logger_file_path"`
	Encoding string `mapstructure:"logger_encoding"`
	Level    string `mapstructure:"logger_level"`
	Logger   string `mapstructure:"logger_logger"`
}

type PostgresConfig struct {
	Host            string        `mapstructure:"postgres_host"`
	Port            string        `mapstructure:"postgres_port"`
	User            string        `mapstructure:"postgres_user"`
	Password        string        `mapstructure:"postgres_password"`
	DbName          string        `mapstructure:"postgres_db_name"`
	SSLMode         string        `mapstructure:"postgres_ssl_mode"`
	MaxIdleConns    int           `mapstructure:"postgres_max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"postgres_max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"postgres_conn_max_lifetime"`
}

type RedisConfig struct {
	Host               string        `mapstructure:"redis_host"`
	Port               string        `mapstructure:"redis_port"`
	Password           string        `mapstructure:"redis_password"`
	Db                 string        `mapstructure:"redis_db"`
	DialTimeout        time.Duration `mapstructure:"redis_dial_timeout"`
	ReadTimeout        time.Duration `mapstructure:"redis_read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"redis_write_timeout"`
	IdleCheckFrequency time.Duration `mapstructure:"redis_idle_check_frequency"`
	PoolSize           int           `mapstructure:"redis_pool_size"`
	PoolTimeout        time.Duration `mapstructure:"redis_pool_timeout"`
}

type PasswordConfig struct {
	IncludeChars     bool `mapstructure:"password_include_chars"`
	IncludeDigits    bool `mapstructure:"password_include_digits"`
	MinLength        int  `mapstructure:"password_min_length"`
	MaxLength        int  `mapstructure:"password_max_length"`
	IncludeUppercase bool `mapstructure:"password_include_uppercase"`
	IncludeLowercase bool `mapstructure:"password_include_lowercase"`
}

type CorsConfig struct {
	AllowOrigins string `mapstructure:"cors_allow_origins"`
}

type OtpConfig struct {
	ExpireTime time.Duration `mapstructure:"otp_expire_time"`
	Digits     int           `mapstructure:"otp_digits"`
	Limiter    time.Duration `mapstructure:"otp_limiter"`
}

type JWTConfig struct {
	AccessTokenExpireDuration  time.Duration `mapstructure:"jwt_access_token_expire_duration"`
	RefreshTokenExpireDuration time.Duration `mapstructure:"jwt_refresh_token_expire_duration"`
	Secret                     string        `mapstructure:"jwt_secret"`
	RefreshSecret              string        `mapstructure:"jwt_refresh_secret"`
}

func GetConfig() *Config {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling environment variables: %v", err)
	}

	if port := os.Getenv("PORT"); port != "" {
		cfg.Server.ExternalPort = port
		log.Printf("Overriding external port from PORT env -> %s", port)
	}

	return &cfg
}
