package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App            AppConfig            `yaml:"app"`
	Database       DatabaseConfig       `yaml:"database"`
	JWT            JWTConfig            `yaml:"jwt"`
	Cloudinary     CloudinaryConfig     `yaml:"cloudinary"`
	Redis          RedisConfig          `yaml:"redis"`
	MeiliSearch    MeiliSearch          `yaml:"meilisearch"`
	MongoDB        MongoDBConfig        `yaml:"mongodb"`
	PaymentGateway PaymentGatewayConfig `yaml:"payment_gateway"`
}

type PaymentGatewayConfig struct {
	SecretKey          string `yaml:"secret_key"`
	SuccessRedirectUrl string `yaml:"success_redirect_url"`
	FailureRedirectUrl string `yaml:"failure_redirect_url"`
}

type MongoDBConfig struct {
	URI            string `yaml:"uri"`
	DatabaseName   string `yaml:"database_name"`
	CollectionName string `yaml:"collection_name"`
}

type MeiliSearch struct {
	Username string `yaml:"username"`
	APIKey   string `yaml:"api_key"`
}

type CloudinaryConfig struct {
	Name      string `yaml:"name"`
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

type RedisConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type JWTConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	User string `yaml:"db_user"`
	Pass string `yaml:"db_pass"`
	Host string `yaml:"db_host"`
	Port string `yaml:"db_port"`
	Name string `yaml:"db_name"`
}

func LoadConfig(filename string) Config {
	var Cfg Config
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		return Cfg
	}

	err = yaml.Unmarshal(fileByte, &Cfg)
	if err != nil {
		panic(err)
	}
	return Cfg
}
