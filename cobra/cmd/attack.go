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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// attackCmd represents the attack command
var attackCmd = &cobra.Command{
	Use:   "attack [service]",
	Short: "A brief description of your command",
	Long:  `A longer description that ....`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := make(map[string]string)
		m["service1"] = "service1 called"
		m["service2"] = "service2 called"
		m["service3"] = "service3 called"

		msg, ok := m[args[0]]
		if !ok {
			_, _ = fmt.Fprintln(os.Stderr, "unknown service")
			return
		}

		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(attackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
