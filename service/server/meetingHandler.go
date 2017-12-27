package server

import (
	"database/sql"
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"time"
	"github.com/unrolled/render"
	"agenda_api/cli/service"
	"agenda_api/cli/entity"
)

func createMeetingHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		successMsg.Message = "creat meeting success"
		failMsg.Message = "create meeting fail"
		
		//check if login
		cookie, err := req.Cookie("LoginId")
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if err1 != nil{
			failMsg.Data.Error = err1.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		
		//get meeting info
		var reqBody struct {
			Title string
			Participants []string
			Starttime time.Time
			Endtime time.Time
		}
		if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil{
			failMsg.Data.Error = "[creatMeetingHandler][DecodeRequest]:"+err.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		fmt.Println(reqBody.Participants)
		//check time conflict
		relativePeople := reqBody.Participants
		relativePeople = append(relativePeople, session.CurrentUser)
		for _,person := range relativePeople {
			User, err := service.UserInfoService.FindByUsername(person)
			if err != nil || User == nil {
				failMsg.Data.Error = "[creatMeetingHandler]:" + person + ":no such user" 
				formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
				return
			}
			Meetings, _ := service.MeetingInfoService.GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone(person, reqBody.Starttime, reqBody.Endtime)
			if Meetings != nil && len(Meetings) != 0 {
				failMsg.Data.Error = "[creatMeetingHandler]:" + person + "have time conflict" 
				formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
				return
			}
		}
		meeting := entity.Meeting{
			Title:reqBody.Title, 
			Sponsor:session.CurrentUser, 
			Participants:sql.NullString{Valid:true, String:strings.Join(reqBody.Participants, "&")},
			StartTime:reqBody.Starttime,
			EndTime:reqBody.Endtime,
		}
		if err := service.MeetingInfoService.InsertMeeting(meeting); err != nil {
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		formatter.JSON(w, http.StatusOK, successMsg)
	}
}

func addParticipatorsHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		
	}
}

func deleteParticipatorsHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		
	}
}

func cncelMeetingHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		successMsg.Message = "cancle a meeting success"
		failMsg.Message = "cancle a meeting fail"
		
		//check if login
		cookie, err := req.Cookie("LoginId")
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if err1 != nil{
			failMsg.Data.Error = err1.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}

		//get the title name and delete it
		path := req.URL.Path
		pathSlice := strings.Split(path, "/")
		title :=  pathSlice[len(pathSlice) - 1]
		//if it's delete all meeting
		if title == "all" {
			fmt.Println("enter {title}")
			return;
		}
		fmt.Println(title)
		meeting, err3:= service.MeetingInfoService.FindByTitle(title)
		if err3 != nil {
			failMsg.Data.Error = err3.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		if meeting.Sponsor != session.CurrentUser {
			failMsg.Data.Error = "The meeting's Sponsor is not you!"
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		err = service.MeetingInfoService.DeleteAMeetingByTitle(title)
		if err != nil {
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		formatter.JSON(w, http.StatusOK, successMsg)
	}
}

func quitMeetingHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		successMsg.Message = "quit a meeting success"
		failMsg.Message = "quit a meeting fail"
		
		//check if login
		cookie, err := req.Cookie("LoginId")
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if err1 != nil{
			failMsg.Data.Error = err1.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}

		//get the meeting by title
		path := req.URL.Path
		pathSlice := strings.Split(path, "/")
		title :=  pathSlice[len(pathSlice) - 1]
		meeting, err3:= service.MeetingInfoService.FindByTitle(title)
		if err3 != nil {
			failMsg.Data.Error = err3.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		
		//check if you are participants of the meeting
		participants := meeting.Participants.String
		if !strings.Contains(participants, session.CurrentUser) {
			failMsg.Data.Error = "you are not participants of the meeting!"
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		participantSlice := strings.Split(participants, "&")
		resultParticipantsSlice := make([]string, 0)
		for _, person := range participantSlice {
			if person != session.CurrentUser {
				resultParticipantsSlice = append(resultParticipantsSlice, person)
			}
		}
		resultParticipantsString := strings.Join(resultParticipantsSlice, "&")
		err = service.MeetingInfoService.UpdateMeetingParticipants(title, resultParticipantsString)
		if err != nil {
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		formatter.JSON(w, http.StatusOK, successMsg)
	}
}

func deleteAllMeetingHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		fmt.Println("enter {all}")
		successMsg.Message = "cancle sponsor meetings success"
		failMsg.Message = "cancle sponsor meetings fail"
		
		//check if login
		cookie, err := req.Cookie("LoginId")
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if err1 != nil{
			failMsg.Data.Error = err1.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}

		//delete all sponsor meetings
		err = service.MeetingInfoService.DeleteAllSponsorMeeting(session.CurrentUser)
		if err != nil {
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		formatter.JSON(w, http.StatusOK, successMsg)
	
	}
}

func queryMeetingHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		var querySuccessMsg struct {
			Message string
			Data struct{
				MeetingList []entity.Meeting
			}
		}
		querySuccessMsg.Message = "query meetings success"
		failMsg.Message = "query meetings fail"
		
		//check if login
		cookie, err := req.Cookie("LoginId")
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}
		session,err1 := service.UserInfoService.UserHasLogin(cookie.Value)
		if err1 != nil{
			failMsg.Data.Error = err1.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		}

		//query the meeting
		var reqBody struct {
			StartTime string
			EndTime string
		}
	//	fmt.Println(reqBody.StartTime, reqBody.EndTime)
		if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil{
			failMsg.Data.Error = "[queryMeetingHandler][DecodeRequest]:"+err.Error()
			formatter.JSON(w, http.StatusUnprocessableEntity, failMsg)
			return
		}
		fmt.Println(reqBody.StartTime, reqBody.EndTime)
		startTime, _ := String_to_date(reqBody.StartTime + "/00:00")
		endTime, _ :=  String_to_date(reqBody.EndTime + "/00:00")
		fmt.Println(startTime, endTime)
		fmt.Println(session.CurrentUser)
		meetingList, err := service.MeetingInfoService.GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone(session.CurrentUser, startTime, endTime)
		fmt.Println(meetingList)
		if err != nil{
			failMsg.Data.Error = err.Error()
			formatter.JSON(w, http.StatusBadRequest, failMsg)
			return
		} else {
			querySuccessMsg.Data.MeetingList = meetingList
			formatter.JSON(w, http.StatusOK, querySuccessMsg)
		}
	}
}


func Date_to_string(date time.Time) string {
	return date.Format("2006-01-02/15:04")
}

func String_to_date(date string) (time.Time, error) {
	the_time, err := time.Parse("2006-01-02/15:04", date)
	return the_time, err
}