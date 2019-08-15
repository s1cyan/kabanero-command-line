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
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var kabConfig *viper.Viper

func urlAccess(url string) {
	if url == "" {
		fmt.Printf("URLLLLLLLLLL?>>>>>>>> + " + kabConfig.GetString("KABURL"))
	} else {
		kabConfig.SetDefault("KABURL", url)
	}
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(2),
	Use:   "login userid password kabanero-url",
	Short: "Will authentic you to the Kabanero instance",
	Long: `
	The userid and password passed will be used
	to authenticate the user with kabanero instance.
	
	By authenticating with the Kabanero instance, 
	you will be able to manage the instance of kabanero.`,
	Example: `
		kabanero-management champ champpassword https://kabanero1.io
		`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("login called")

		username := args[0]
		password := args[1]
		var kabURL string
		KabEnvVar := "KABURL"

		if len(args) > 2 {
			kabURL = args[2]
			os.Setenv(KabEnvVar, kabURL)
			urlAccess(kabURL)
			// fmt.Printf("SET VAR?----" + os.Getenv(KabEnvVar))
			// cliConfig.SetDefault(KabEnvVar, kabURL)
			// fmt.Printf("\n VIPER ACCESS ------" + cliConfig.GetString(KabEnvVar))

		} else {
			// cliConfig.AutomaticEnv()

			val := kabConfig.GetString(KabEnvVar)
			// val, present := os.LookupEnv("KABURL_MANAGEMENTCLI")
			urlAccess("")
			if val == "" {
				return errors.New("No Kabanero instance url specified")
			}
			kabURL = val
		}

		// client := &http.Client{
		// 	Timeout: time.Second * 30,
		// }

		// requestBody, _ := json.Marshal(map[string]string{"gituser": username, "gitpat": password})

		// req, err := http.NewRequest("POST", kabURL, bytes.NewBuffer(requestBody))
		// if err != nil {
		// 	return err
		// }

		// req.Header.Set("Content-Type", "application/json")

		// resp, err := client.Do(req)

		// if err != nil {
		// 	fmt.Printf("The HTTP request failed with error %s\n", err)
		// }

		// defer resp.Body.Close()

		// data := json.Decode(resp.Body)

		// fmt.Println(string(data))
		fmt.Printf("USERNAME/PWD/KAB" + username + "-- " + password + "____" + kabURL)
		return nil
	},
}

func init() {
	kabConfig = viper.New()
	if cfgFile != "" {
		kabConfig.SetConfigName("kabConfig")
		kabConfig.SetConfigType("yaml")

	}
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
