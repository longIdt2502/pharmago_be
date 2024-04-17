package utils

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	EmailSenderName     string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	B2KeyId             string        `mapstructure:"B2_KEY_ID"`
	B2KeyName           string        `mapstructure:"B2_KEY_NAME"`
	B2ApplicationKey    string        `mapstructure:"B2_APPLICATION_KEY"`
	B2Bucket            string        `mapstructure:"B2_BUCKET"`
	B2AccountId         string        `mapstructure:"B2_ACCOUNT_ID"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	WezoloServerAdress  string        `mapstructure:"WEZOLO_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		timeToken, _ := time.ParseDuration("24h15m")
		config.DBDriver = os.Getenv("DB_DRIVER")
		config.DBSource = os.Getenv("DB_SOURCE")
		config.HTTPServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
		config.GRPCServerAddress = os.Getenv("GRPC_SERVER_ADDRESS")
		config.RedisAddress = os.Getenv("REDIS_ADDRESS")
		config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
		config.AccessTokenDuration = timeToken
		config.EmailSenderName = os.Getenv("EMAIL_SENDER_NAME")
		config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
		config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")
		config.B2KeyId = os.Getenv("B2_KEY_ID")
		config.B2KeyName = os.Getenv("B2_KEY_NAME")
		config.B2ApplicationKey = os.Getenv("B2_APPLICATION_KEY")
		config.B2Bucket = os.Getenv("B2_BUCKET")
		config.B2AccountId = os.Getenv("B2_ACCOUNT_ID")
		config.MigrationURL = os.Getenv("MIGRATION_URL")
		config.MigrationURL = os.Getenv("WEZOLO_SERVER_ADDRESS")

		err = nil

		return
	}

	err = viper.Unmarshal(&config)
	return
}
