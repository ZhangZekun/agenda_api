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
//	"agenda_api/cli/entity"
	"os"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"strings"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in the agenda",
	Long: `you need to input the username and password,for example:
	./agenda login -u=zhangzemian -p=12345678`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		//set request
		client := &http.Client{}
		reqBody := "{"+"\"Username\":\"" + username +"\", \"Password\":\"" + password + "\"}"
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/user/login", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("err when create the login request")
		}
		cookieId := GetCookieID()
		if cookieId != "" {
			req.Header.Set("Cookie", "LoginId="+cookieId)
		}

		//send request and output response body
		resp, err := SendRequestAndOutputResponseBody(req, client)
		if err != nil {
			fmt.Println(err.Error());
		}

		//save cookies pass by the server
		if err:=saveCookies(resp); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "User password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func saveCookies(resp *http.Response) error {
	cookies := resp.Cookies()
	if len(cookies) != 0 {
		byteData, _ := json.Marshal(cookies)
		//写入文件
		file, error := os.Create("cookies.txt");
		defer file.Close()
		if error != nil {
			return error
		}
		file.Write(byteData)
	}
	return nil
}