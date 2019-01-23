// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cvgw/rds_provider/pkg/provider/subnet_group"
	"github.com/spf13/cobra"
)

// createSubnetGroupCmd represents the createSubnetGroup command
var createSubnetGroupCmd = &cobra.Command{
	Use:   "createSubnetGroup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createSubnetGroup called")

		fileData, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		req := subnet_group.CreateSubnetGroupRequest{}
		err = json.Unmarshal(fileData, &req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(req)

		input := subnet_group.NewCreateDBSubnetGroupInput(req)

		fmt.Println(input)
	},
}

var (
	file string
)

func init() {
	rootCmd.AddCommand(createSubnetGroupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createSubnetGroupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createSubnetGroupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createSubnetGroupCmd.PersistentFlags().StringVarP(
		&file, "file", "f", "", "input file for the resource",
	)
}
