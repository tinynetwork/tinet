// Package cmd tn cli tool
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tinynetwork/tinet/internal/pkg/shell"
	"github.com/tinynetwork/tinet/internal/pkg/utils"

	"github.com/spf13/viper"
)

var defaultCfgFile string = "spec.yaml"
var cfgFile string
var tnconfig shell.Tn

// var verbose bool = false
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tinet",
	Short: "tn: tinet",
	Long:  `tinet is network emulator created by docker`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	checkCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	confCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	downCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	execCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	initCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	imgCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	printCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	psCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	pullCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	reconfCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	reupCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	testCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	upCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")
	upconfCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfgFile, "config file (default is ./spec.yaml)")

	// verbose option
	confCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	upconfCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	upCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	reupCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	reconfCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	execCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")
	downCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose (default: false)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if !utils.Exists(cfgFile) {
		err := fmt.Errorf("%s is not Found", cfgFile)
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&tnconfig); err != nil {
		panic(err)
	}
}
