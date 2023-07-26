/*
Copyright © 2023 Vivek Kumar Sahu  viveksahu650@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "cipher_cli.yaml"

// rootCmd represents the base command when called without any subcommands
func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "viver",
		Short: "Encrypt and decrypt secret messages in seconds!!!",
		Long: `This application encrypts and decrypts secret messages with ease. 

Try out the Caesar and Bacon Cipher options to generate secret messages and share with your inner circle

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=caesar --key=54

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=bacon`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	New().PersistentFlags().StringP("key", "k", "", "The key to pass for the algorithm")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	New().Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra_cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra_cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
