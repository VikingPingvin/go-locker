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
	"vikingPingvin/locker/locker"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
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
		//viper.SetConfigFile(cfgFile)

		cleanenv.ReadConfig(cfgFile, &locker.LockerAgentConfig)
	}

	locker.LockerAgentConfig.Agent.ArgPath = input.Path
	locker.LockerAgentConfig.Agent.ArgNamespace = input.Namespace
	locker.LockerAgentConfig.Agent.ArgConsume = input.Consume

}
