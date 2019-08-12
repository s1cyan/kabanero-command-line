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
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List all the collections in the apphub, and optionally their status",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("list called")

		// var endpoint string = "http://10.211.54.131/api/4.0/edges/edge-3/firewall/config"
		// var endpoint string = "http://10.211.54.244:31000/KabCollections-1.0-SNAPSHOT/v1/collections"
		var endpoint string = "http://9.12.77.136:9080/kabasec/login"

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		client := &http.Client{
			Timeout: time.Second * 30,
		}
		// formData := url.Values{
		// 	"gituser": {username},
		// 	"gitpat":  {password},
		// }

		requestBody, _ := json.Marshal(map[string]string{"gituser": username, "gitpat": password})

		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Printf("Error : %s", err)
		}

		req.Header.Set("Content-Type", "application/json")

		// req.SetBasicAuth(username, password)

		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}

		// data, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		data := json.Decode(resp.Body)
		// data = json.Decode(resp.Body)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringP("username", "u", "", "nsxmanager username")
	listCmd.PersistentFlags().StringP("password", "p", "", "nsxmanagerpassword")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
