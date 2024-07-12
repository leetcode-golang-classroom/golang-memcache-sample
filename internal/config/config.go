package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port          int    `mapstructure:"PORT"`
	MemcacheURL   string `mapstructure:"MEMCACHE_URL"`
	JsonServerURL string `mapstructure:"JSONSERVER_URL"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	util.FailOnError(v.BindEnv("PORT"), "failed to bind env PORT")
	util.FailOnError(v.BindEnv("MEMCACHE_URL"), "failed to bind env MEMCACHE_URL")
	util.FailOnError(v.BindEnv("JSONSERVER_URL"), "failed to bind env JSONSERVER_URL")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("load from environment variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		util.FailOnError(err, "Failed to read enivroment")
	}
}
