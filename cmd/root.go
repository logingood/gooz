// Copyright © 2019 Murat Mukhtarov <https://github.com/logingood/gooz>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile               string
	ticketsFilePath       string
	organizationsFilePath string
	usersFilePath         string
	validInput            map[string]bool
	validStrings          []string
	validInputUsers       map[string]bool
	validStringsUsers     []string
	validInputTickets     map[string]bool
	validStringsTickets   []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gooz",
	Short: "Zendesk Code Challenge (Search CLI)",
	Long: `Enjoy search across three json files using any field from json schema.
	Gooz will draw for you tables with the results.
	`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gooz.yaml)")
	rootCmd.PersistentFlags().StringVar(&ticketsFilePath, "tickets_path", "data/tickets.json", "path to your tickets.json, default is data/tickets.json")
	rootCmd.PersistentFlags().StringVar(&organizationsFilePath, "organizations_path", "data/organizations.json", "path to your organizations.json, default is data/organizations.json")
	rootCmd.PersistentFlags().StringVar(&usersFilePath, "users_path", "data/users.json", "path to your users.json, default is data/users.json")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gooz" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gooz")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
