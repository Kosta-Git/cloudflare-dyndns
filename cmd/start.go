package cmd

import (
	"github.com/Kosta-Git/cloudflare-dyndns/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Zone    string        `mapstructure:"zone"`
	SubZone string        `mapstructure:"sub_zone"`
	ApiKey  string        `mapstructure:"api_key"`
	Daemon  time.Duration `mapstructure:"daemon"`
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the DynDNS updater",
	Run: func(cmd *cobra.Command, args []string) {
		var configuration Config
		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to unmarshal config, %v", err)
		}
		log.Print("Starting DynDNS updater")
		for {
			pkg.UpdateDynDns(configuration.Zone, configuration.SubZone, configuration.ApiKey)
			time.Sleep(configuration.Daemon)
		}
	},
}

func init() {
	viper.SetConfigName("cf-dyndns")
	viper.AddConfigPath("/etc/cf-dyndns")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v\nConfig file must be either in . or /etc/cf-dyndns, and must be named cf-dyndns\n", err)
	}

	rootCmd.AddCommand(startCmd)
}
