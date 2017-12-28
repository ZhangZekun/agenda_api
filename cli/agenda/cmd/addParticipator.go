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
	"strings"
)

// operateParticipantCmd represents the operateParticipant command
var addParticipatorCmd = &cobra.Command{
	Use:   "addParticipator",
	Short: "the sponsor of the meeting can add  in the meeting",
	Long: `you need to input three arguments(title(t), operation(o), participants(p)).For example:
	addCMD:./agenda operateParticipant -t=Work -o=add -p=zhangzekun, zhangzhijian;
	deleteCMD:./agenda operateParticipant -t=Work -o=del -p=zhangzekun,zhangzhijian`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addParticipator called")
		title, _ := cmd.Flags().GetString("title")
		participator, _:= cmd.Flags().GetString("participants")
		//set request
		client := &http.Client{}
		reqBody := fmt.Sprintf("{\"Username\":\"%s\"}",participator)	
		fmt.Println(reqBody)
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/meeting/"+title +"/participators", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("err when create creatMeeting request")
		}
		cookieId := GetCookieID()
		if cookieId != "" {
			req.Header.Set("Cookie", "LoginId="+cookieId)
		}

		//send request and output response body
		_, err = SendRequestAndOutputResponseBody(req, client)
		if err != nil {
			fmt.Println(err.Error());
		}
	},
}

func init() {
	RootCmd.AddCommand(addParticipatorCmd)
	addParticipatorCmd.Flags().StringP("title", "t", "", "title")
	addParticipatorCmd.Flags().StringP("participants", "p", "", "Names of Participants")



	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// operateParticipantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// operateParticipantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
