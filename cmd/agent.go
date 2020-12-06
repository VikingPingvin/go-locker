/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"vikingPingvin/locker/locker"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Input struct {
	Path      string
	Namespace string
	Consume   string
}

// Pointer to an Input Instance
var input *Input

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start Locker in Agent mode",
	Long:  `locker agent`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		locker.ExecuteAgent()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Initialize input to avoid nil pointer dereference errors
	input = &Input{}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	agentCmd.Flags().StringVar(&input.Path, "file", "", "[path,...] Absolute or relative path(s). Multiple paths must be separated with ','")
	agentCmd.Flags().StringVar(&input.Namespace, "namespace", "", "[namespace/project/job-id] Separator must be '/'")
	agentCmd.Flags().StringVar(&input.Consume, "consume", "", "[namespace/project/job-id] Requests the specified artifact to download from the Locker Server.")

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

		// Search config in home directory with name ".locker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".locker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}

	cfgContent = viper.AllSettings()

	log.Debug().Msgf("config:\n %v", viper.AllSettings())

	locker.LockerConfig = &locker.AgentConfig{}
	locker.LockerConfig.ServerIP = viper.GetString("agentconfig.server_ip")
	locker.LockerConfig.ServerPort = viper.GetString("agentconfig.server_port")
	locker.LockerConfig.LogPath = viper.GetString("agentconfig.log_path")
	locker.LockerConfig.SendConcurrent = viper.GetBool("agentconfig.send_concurrent")

	locker.LockerConfig.ArgPath = input.Path
	locker.LockerConfig.ArgNamespace = input.Namespace
	locker.LockerConfig.ArgConsume = input.Consume
}
