// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"net/http"
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "user log out the agenda",
	Long: `you don't need to input anything.For example:
	./agenda logout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		//set request
		client := &http.Client{}
		
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/user/logout", nil)
		if err != nil {
			fmt.Println("err when create the logout request")
			return
		}
		cookieId := GetCookieID()
		if cookieId != "" {
			req.Header.Set("Cookie", "LoginId="+cookieId)
		}

		//send request and output response body
		_, err = SendRequestAndOutputResponseBody(req, client)
		if err != nil {
			fmt.Println(err.Error());
			return
		}

		//clear Cookies
		clearCookies()
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func clearCookies() {
	file, error := os.Create("cookies.txt");
	defer file.Close()
	if error != nil {
		return
	}
	file.Write([]byte(""))
}