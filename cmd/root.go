package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string
var registry string
var json bool
var path string
var version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gomodctl",
	Short:   "Check and Update Go modules.",
	Long:    `gomodctl is a Go tool that provides check and update features for Go modules.`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	debug.SetGCPercent(-1)

	ctx, cncl := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cncl()

	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/gomodctl.yml)")
	rootCmd.PersistentFlags().StringVar(&registry, "registry", "", "URI of the registry to be used for search")
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "Print JSON result")
	rootCmd.PersistentFlags().StringVar(&path, "path", "", "Optional go.mod parent directory")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("registry", rootCmd.PersistentFlags().Lookup("registry"))
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")

	if config != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if path != "" {
			viper.AddConfigPath(path)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)

		viper.SetConfigName("gomodctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println(err)
	}
}
