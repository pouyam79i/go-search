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
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "This command is used to search within a file",
	Long: `Use this command to search for a specific data in a file or analize data in a file. Also you can save the result of searching!

This command is written in go lang and by using cobra pkg.
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args)
		textStatus, _ := cmd.Flags().GetBool("text")
		if textStatus {
			data, err := analizeTextFile(args[0])
			if err != nil {
				fmt.Println("Failed to analize text file")
			} else {
				fmt.Println(data)
				saveState, _ := cmd.Flags().GetBool("save")
				if saveState {
					saveResult(data, args[1])
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	// Used for analizing text file
	fileCmd.Flags().BoolP("text", "t", false, "Determines that the file is a text file")
	// Used to save result
	fileCmd.Flags().BoolP("save", "s", false, "Saves the result in the told text file")
}

// Used fucntion for this command:

/**
* This function analizes text files
* fileAddress is the address of text file
* it returns data and error!
* if error is nill it means the process is done successfully
 */
func analizeTextFile(fileAddress string) (string, error) {

	// Opening file
	if strings.TrimSpace(fileAddress) == "" {
		fmt.Printf("Empty file address\n")
		return "NULL DATA", errors.New("empty file address")
	}
	file, err := os.Open(fileAddress)
	if err != nil {
		fmt.Printf("Failed to read '%s' text file\n", fileAddress)
		return "NULL DATA", err
	}
	// fmt.Println("Data founded")
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Print("Failed to colse file")
		}
	}()

	// Reading texts from file
	scanner := bufio.NewScanner(file)
	allStrings := ""
	data := ""
	for scanner.Scan() {
		allStrings += "\n" + scanner.Text()
	}
	// fmt.Println(allStrings)

	// Analizing text file
	wordList := strings.Fields(allStrings)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	for index, element := range counts {
		data += index + " => " + strconv.Itoa(element) + "\n"
	}

	// returning result
	return data, nil
}

func saveResult(data, fileAddress string) error {

	// Creating a new text file to save data
	if strings.TrimSpace(fileAddress) == "" {
		fmt.Printf("Empty file address\n")
		return errors.New("empty file address")
	}
	file, err := os.Create(fileAddress)
	if err != nil {
		fmt.Printf("Failed to create '%s' file.\n result is not saved\n", fileAddress)
	}
	defer file.Close()

	// Writing results in text file
	_, err2 := file.WriteString(data)
	if err2 != nil {
		fmt.Printf("Failed to write result in '%s'\n", fileAddress)
	} else {
		fmt.Println("The result is successfully saved!")
	}

	// Returning results
	return err
}
