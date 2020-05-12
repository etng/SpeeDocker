package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// initConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string) *viper.Viper {
	cfg := viper.New()
	cfg.SetDefault("client.server_url", "http://103.103.49.194:50001")
	cfg.SetDefault("server.port", 50001)
	cfg.SetDefault("server.docker.username", "")
	cfg.SetDefault("server.docker.password", "")

	if cfgFile != "" {
		// Use config file from the flag.
		cfg.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
			//os.Exit(1)
		}

		cfg.AddConfigPath(home)
		if me, err := os.Executable();err==nil{
			viper.AddConfigPath(filepath.Dir(me))
		}
		cfg.SetConfigName("speedocker")
	}

	cfg.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := cfg.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", cfg.ConfigFileUsed())
	}
	return cfg
}
