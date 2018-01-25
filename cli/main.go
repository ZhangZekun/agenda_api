package main

import(
//	"agenda_api/cli/entity"
	"github.com/ZhangZeMian/agenda_api/cli/service"
	"fmt"
	"time"
)

func main() {
//	var err error

	//  testb, err := service.UserInfoService.FindByUsername("test_user1")
	//  if err != nil {
	// 	fmt.Println(err)
	//  }
	// fmt.Println(testb)

	// insertUser := database.User{Username:"zhangzhijian", Password:"555"}
	// err = database.UserInfoService.InsertUser(insertUser)
	// fmt.Println(err)

	// err = database.UserInfoService.LoginInfoInsert("zhangzemian")
	// fmt.Println(err)

	// Boola := database.UserInfoService.UserHasLogin("zhangzhendong")
	// fmt.Println(Boola)

	// err = database.UserInfoService.LoginInfoDelete("zhangzhendong")
	// fmt.Println(err)

	// userlist, err2 := database.UserInfoService.GetAllUsersInfo("zhangzhendong")
	// fmt.Println(err2)
	// fmt.Println(userlist)

	// insertMeeting := database.Meeting{Title:"huiyi", Sponsor:"zhangzekun", StartTime:time.Now(), EndTime:time.Now()}
	// err := database.MeetingInfoService.InsertMeeting("zhangzekun", insertMeeting)
	// fmt.Println(err)


// 	testb, err := service.MeetingInfoService.FindByID("1")
// 	if err != nil {
// 	   fmt.Println(err)
// 	}
//    fmt.Println(*testb)
	
// 	idSlice := make([]string, 0)
// 	idSlice = append(idSlice, "1")
// 	idSlice = append(idSlice, "2")
// 	idSlice = append(idSlice, "3")
// 	StartTime, _ := String_to_date("2014-02-03/00:00")
// 	EndTime, _ := String_to_date("2016-02-03/00:00")
// 	fmt.Println(StartTime, EndTime)
// 	meetinglist, err := service.MeetingInfoService.GetAllMeetingBetweenStartTimeAndEndTimeOfIdList(idSlice, StartTime, EndTime)
// 	if err != nil {	
// 	   fmt.Println(err)
// 	}
//    fmt.Println(meetinglist)

	// insertMeeting := entity.Meeting{Title:"huiyi", Sponsor:"zhangzekun", StartTime:time.Now(), EndTime:time.Now()}
	// err := service.MeetingInfoService.InsertMeeting(insertMeeting)
	// fmt.Println(err)

	// err := service.MeetingInfoService.UpdateMeetingParticipants("huiyi4", "1&2&3")
	// fmt.Println(err)

	// err := service.MeetingInfoService.DeleteAMeetingByTitle("huiyi4")
	// fmt.Println(err)

	// err := service.MeetingInfoService.DeleteAllSponsorMeeting("zhangzekun")
	// fmt.Println(err)

	// meetings,err := service.MeetingInfoService.GetAllRelativeMeeting("zhangzemian")
	// fmt.Println(meetings)
	// fmt.Println(err)

	StartTime, _ := String_to_date("2014-02-03/00:00")
	EndTime, _ := String_to_date("2017-02-03/00:00")
	meetings,err := service.MeetingInfoService.GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone("zhangzemian", StartTime, EndTime)
	fmt.Println(meetings)
	fmt.Println(err)
}

func Date_to_string(date time.Time) string {
	return date.Format("2006-01-02/15:04")
}

func String_to_date(date string) (time.Time, error) {
	the_time, err := time.Parse("2006-01-02/15:04", date)
	return the_time, err
}