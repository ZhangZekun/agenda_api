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
	"fmt"
	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "to delete the user own account",
	Long: `you don't need to input anything.For example:
	./agenda deleteUser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteUser called")

		//	fmt.Println("register called by " + username + password + email)

		//set request
		client := &http.Client{}
		
		req, err := http.NewRequest("DELETE", "http://localhost:8080/api/agenda/user/self", nil)
		if err != nil {
			fmt.Println("err when create the delete user request")
			return
		}
		fmt.Println("1");
		cookieId := GetCookieID()
		if cookieId != "" {
			req.Header.Set("Cookie", "LoginId="+cookieId)
		}
		fmt.Println("2");
		//send request and output response body
		_, err = SendRequestAndOutputResponseBody(req, client)
		if err != nil {
			fmt.Println(err.Error() + "aaaa");
			return
		}
		fmt.Println("3");
		//clear Cookies
		clearCookies()
	},
}

func init() {
	RootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
