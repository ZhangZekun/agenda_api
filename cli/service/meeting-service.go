package service
import (
	"time"
	"errors"
	"strings"
	"agenda_api/cli/entity"
	"agenda_api/cli/dao"
)
//UserInfoAtomicService .
type MeetingInfoAtomicService struct{}

//UserInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}




//InsertAUser
func (*MeetingInfoAtomicService) InsertMeeting(session entity.Session, meeting entity.Meeting) error {
	//check if log in 
	if login:=UserInfoService.UserHasLogin(session.LoginId); login==false {
		err:= errors.New("You haven't log in!")
		return err
	}
	//check if the user exist and get the user's infomation
	// var UserInfo *User
	// var err error
	// if UserInfo, err = UserInfoService.FindByUsername(username); err != nil {
	// 	return err
	// }
	// UsersNamePaticipantTheMeeting := strings.Split(meeting.Participants.String, "&")
	//
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.MeetingInfoDao{tx}
    err = dao.InsertMeeting(session.CurrentUser, meeting)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//query all meeting infomation betwwen startTime and endTime for login user
//this is used for http server!
func (*MeetingInfoAtomicService) GetAllMeetingBetweenStartTimeAndEndTimeOfLoginUser(session entity.Session , startTime time.Time, endTime time.Time) ([]entity.Meeting, error) {
	if login:=UserInfoService.UserHasLogin(session.LoginId); login==false {
		err:= errors.New("You haven't log in!")
		return nil, err
	}
	return nil, nil
}

//query all meeting infomation betwwen startTime and endTime for any user
//this is used for this package!

func (*MeetingInfoAtomicService) GetAllMeetingBetweenStartTimeAndEndTimeOfAnyUser(username string , startTime time.Time, endTime time.Time) ([]entity.Meeting, error) {
	//if no such user in database
	var UserInfo *entity.User
	var err error
	if UserInfo, err = UserInfoService.FindByUsername(username); err != nil {
		return nil, err
	}
	allRelativeMeetingIDArray := strings.Split(UserInfo.SponsorMeeting.String+"&"+UserInfo.ParticipantMeeting.String, "&")
	return nil, nil
}
