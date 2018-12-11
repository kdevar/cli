// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Data struct {
	Name string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create project from template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		done := make(chan struct{}) // HLdone
		defer close(done)
		c := fileProcesser(done, "templates/src/")

		for template := range c {
			fmt.Println(template)
		}

	},
}

func fileProcesser(done <-chan struct{}, root string) <- chan string{
	files := make(chan string, 15)
	results := make(chan string, 15)

	go func(){
		err := filepath.Walk(root, GetFiles(files))

		if err != nil {
			close(files)
		}
	}()
	return results
}

func GetFiles(c chan string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {
		fmt.Println("file walker", path)
		c <- path
		return nil
	}
}

func processFiles(path string, r chan string) {
	fmt.Println("process", path)
	r <- path
}

func init() {
	rootCmd.AddCommand(initCmd)
}
