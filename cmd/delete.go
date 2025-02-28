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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To Delete an AKS cluster",
	Long:  `You can delete a cluster with aksctl delete cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Red("delete is not in aksctl command group. See 'aksctl delete --help' ")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
