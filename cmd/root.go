/*
Copyright © 2020 Søren Nielsen <contact@cph.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/sorennielsen/bmi/internal/bmi"
	"github.com/spf13/cobra"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var arg0 = determineAppNameFromArg0()
var shortDesc = `Calculate BMI from height and weight either directly 
from the command line or by using the built-in web service.`
var longDesc = shortDesc + `

Example: 

	$ ` + arg0 + ` 186 85

Output: ` + example + `

Or start the web service with:

	$ ` + arg0 + ` serve

`

// Determine configuration name from arg0
// This means that this value will change if the program is
// invoked through a symbolic link fx.
func determineAppNameFromArg0() string {
	arg0 := os.Args[0]
	arg0 = path.Base(arg0)
	return arg0
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   arg0,
	Short: shortDesc,
	Long:  longDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			cmd.Usage()
			return errors.New(`Expect either the 'serve' command or two arguments: height (cm) and weight (kg).`)
		}
		return nil
	},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, desc, err := bmi.Calculate(args[0], args[1])
		if err != nil {
			return err
		}
		fmt.Println(desc)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bmi.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bmi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bmi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
