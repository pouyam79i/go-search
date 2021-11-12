/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// Edited by Pouya Mohammadi 

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "You can see all sub command available:",
	Long: `This program can search for specific data in a file.
This program is written in go lang using cobra pkg.
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		helpMassage := `This program can search for specific data in a file:
	file	If you want to search within a file use this command
	help 	Shows all information about usage of search command
	
	-t, --text   Determines that the file is a text file
	-s, --save   Saves the result in the told text file
	`
		fmt.Println(helpMassage)
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
