package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server  ServerConfig  `mapstructure:"server" json:"server"`
	Logger  Logger        `mapstructure:"logger" json:"logger"`
	MongoDB MongodbConfig `mapstructure:"mongodb" json:"mongodb"`
}

type MongodbConfig struct {
	DatabaseName   string `mapstructure:"db_name" json:"db_name"`
	DatabaseHosts  string `mapstructure:"hosts" json:"hosts"`
	TimeOut        int    `mapstructure:"timeout" json:"timeout"`
	DialTimeOut    int64  `mapstructure:"dial_timeout" json:"dial_timeout"`
	PoolSize       int    `mapstructure:"pool_size" json:"pool_size"`
	Username       string `mapstructure:"username" json:"username"`
	Password       string `mapstructure:"password" json:"password"`
	ReplicaSet     string `mapstructure:"replica_set" json:"replica_set"`
	AuthSource     string `mapstructure:"auth_source" json:"auth_source"`
	URI            string `mapstructure:"uri" json:"uri"`
	CustomIDPrefix string `mapstructure:"custom_id_prefix" json:"custom_id_prefix"`
}

// Server config struct
type ServerConfig struct {
	Application string `mapstructure:"application" json:"application"`
	AppVersion  string `mapstructure:"appVersion" json:"appVersion"`
	Port        int    `mapstructure:"PORT" json:"port"`
}

// Logger config
type Logger struct {
	Level      string `mapstructure:"level" json:"level"`
	MaxSize    int    `mapstructure:"maxSize" json:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups" json:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge" json:"maxAge"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	for _, k := range viper.AllKeys() {
		fmt.Println("Key ", k)
	}
	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
