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
	"vikingPingvin/locker/locker"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start locker in server mode",
	Long:  `server mode`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfigServer()
		locker.ExecuteServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func initConfigServer() {
	var serverConfigStruct locker.ServerConfig

	if cfgFile != "" {
		// Use config file from the flag.
		cleanenv.ReadConfig(cfgFile, &locker.LockerServerConfig)
	} else {
		cleanenv.ReadEnv(&serverConfigStruct)
		locker.LockerServerConfig = &serverConfigStruct
	}

	//test := &locker.LockerServerConfig

	fmt.Println(serverConfigStruct)

	//locker.LockerServerConfig = &locker.ServerConfig{}
	//locker.LockerServerConfig.Server.ServerIP = viper.GetString("serverconfig.server_ip")
	//locker.LockerServerConfig.Server.ServerPort = viper.GetString("serverconfig.server_port")
	//locker.LockerServerConfig.Server.ArtifactRootDir = viper.GetString("serverconfig.artifacts_root_dir")
	//locker.LockerServerConfig.Server.LogPath = viper.GetString("serverconfig.log_path")

}
