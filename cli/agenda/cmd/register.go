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
	"strings"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register an account for the your agenda",
	Long: `you need to input user message to create an account, the arguments include usernam, password, email, telephone.For example:
	./agenda register -u=zhangzemian -p=12345678 -e=1106066690@qq.com -t=15018377821`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		usernameCmd, _ := cmd.Flags().GetString("username")
		passwordCmd, _ := cmd.Flags().GetString("password")
		emailCmd, _ := cmd.Flags().GetString("email")
		phoneCmd, _ := cmd.Flags().GetString("telephone")
		//	fmt.Println("register called by " + username + password + email)

		//set request
		client := &http.Client{}
		reqBody := fmt.Sprintf("{\"Username\":\"%s\", \"Password\":\"%s\", \"Email\":\"%s\", \"Phone\":\"%s\"}", usernameCmd, passwordCmd, emailCmd, phoneCmd)		
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/user/register", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("err when create the login request")
		}
		//send request and output response body
		_, err2 := SendRequestAndOutputResponseBody(req, client)
		if err2 != nil {
			fmt.Println(err2.Error())
			return
		}

	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "User password")
	registerCmd.Flags().StringP("email", "e", "", "User email")
	registerCmd.Flags().StringP("telephone", "t", "", "User telephone")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
