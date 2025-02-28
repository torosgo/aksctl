/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var unInstallViper = viper.New()

var unInstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "To install add-on",
	Long:  `addon message`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	addOnCmd.AddCommand(unInstallCmd)

}