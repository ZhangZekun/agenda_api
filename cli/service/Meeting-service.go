package service
import (
	 "time"
	// "errors"
	// "strings"
	"github.com/ZhangZeMian/agenda_api/cli/entity"
    "github.com/ZhangZeMian/agenda_api/cli/dao"
)
//UserInfoAtomicService .
type MeetingInfoAtomicService struct{}

//UserInfoService .
var MeetingInfoService = MeetingInfoAtomicService{}

//insert a meeting
func (*MeetingInfoAtomicService) InsertMeeting(meeting entity.Meeting) error {
	tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.MeetingInfoDao{tx}
    err = dao.InsertMeeting(meeting)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//delete a meeting by title
func ( *MeetingInfoAtomicService) DeleteAMeetingByTitle(title string ) error {
	tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.MeetingInfoDao{tx}
    err = dao.DeleteAMeetingByTitle(title)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//delete some's all sponsor meeting
func ( *MeetingInfoAtomicService) DeleteAllSponsorMeeting(name string ) error {
	tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.MeetingInfoDao{tx}
    err = dao.DeleteAllSponsorMeeting(name)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}


//update meeting participants
func (*MeetingInfoAtomicService) UpdateMeetingParticipants(title, Participants string) error {
	tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.MeetingInfoDao{tx}
    err = dao.UpdateMeetingParticipants(title, Participants)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//get all relative meetings
func (*MeetingInfoAtomicService) GetAllRelativeMeeting(name string) ([]entity.Meeting, error) {
    dao := dao.MeetingInfoDao{entity.Mydb}
    return dao.GetAllRelativeMeeting(name)
}

//get all participants meetings
func (*MeetingInfoAtomicService) GetAllParticipantsMeeting(name string) ([]entity.Meeting, error) {
    dao := dao.MeetingInfoDao{entity.Mydb}
    return dao.GetAllParticipantsMeeting(name)
}

//get all relative meetings between startime and endtime
func (*MeetingInfoAtomicService) GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone(name string, startTime time.Time, endTime time.Time) ([]entity.Meeting, error) {
    dao := dao.MeetingInfoDao{entity.Mydb}
    return dao.GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone(name, startTime, endTime)
}

// Find a meeting By Meeting Id.
func (*MeetingInfoAtomicService) FindByID(id string) (*entity.Meeting, error) {
    dao := dao.MeetingInfoDao{entity.Mydb}
    return dao.GetMeetingByID(id)
}

// Find a meeting By Meeting Title.
func (*MeetingInfoAtomicService) FindByTitle(title string) (*entity.Meeting, error) {
    dao := dao.MeetingInfoDao{entity.Mydb}
    return dao.GetMeetingByTitle(title)
}

// // Find meetings by a list of id and a startTime and endTime
// func (*MeetingInfoAtomicService) GetAllMeetingBetweenStartTimeAndEndTimeOfIdList(idSlice []string, startTime time.Time, endTime time.Time) ([]entity.Meeting, error) {
//     dao := dao.MeetingInfoDao{entity.Mydb}
//     return dao.GetAllMeetingBetweenStartTimeAndEndTimeOfIdList(idSlice, startTime, endTime)
// }