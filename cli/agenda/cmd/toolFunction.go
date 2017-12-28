
package cmd

import (
//	"agenda_api/cli/entity"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
//get the cookie id, if exist,return id;else return ""
func GetCookieID() string {
	file1, error1 := os.Open("cookies.txt");
	if error1 != nil {
		fmt.Println(error1);
	}
	defer file1.Close();
	buf := make([]byte, 4024);
	byteNum, err1 := file1.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return ""
	}
	// var cookieSlice interface{}
	var cookieSlice []http.Cookie
	if (byteNum != 0) {
		err := json.Unmarshal(buf[0:byteNum], &cookieSlice)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
	}
	
	for _, cookie := range cookieSlice {
		if cookie.Name == "LoginId" && cookie.Path == "/api/agenda/" {
			return cookie.Value
		}
	}
	return ""
}

func SendRequestAndOutputResponseBody(req *http.Request, client *http.Client) (*http.Response, error){
	resp, err := client.Do(req)
	fmt.Println(resp.Header)
	if err != nil {
		return nil, err
	}

	//output the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return resp, nil

}