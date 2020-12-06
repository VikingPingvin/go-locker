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

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start locker in server mode",
	Long:  `server mode`,
	Run: func(cmd *cobra.Command, args []string) {
		locker.ExecuteServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	initConfigServer()
}

func initConfigServer() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}

	locker.LockerServerConfig = &locker.ServerConfig{}
	locker.LockerServerConfig.ServerIP = viper.GetString("serverconfig.server_ip")
	locker.LockerServerConfig.ServerPort = viper.GetString("serverconfig.server_port")
	locker.LockerServerConfig.ArtifactRootDir = viper.GetString("serverconfig.artifacts_root_dir")
	locker.LockerServerConfig.LogPath = viper.GetString("serverconfig.log_path")

}
