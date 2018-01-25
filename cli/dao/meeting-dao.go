package dao
import (
	"fmt"
//		"errors"
		"time"
		"github.com/ZhangZeMian/agenda_api/cli/entity"
	)

type MeetingInfoDao entity.DaoSource

//Insert a Meeting
var meetingInsertStmt = "Insert into Meeting (Title, Sponsor, Participants, StartTime, EndTime) values (?,?,?,?,?)"
func (dao *MeetingInfoDao) InsertMeeting(meeting entity.Meeting) error {
    stmt, err := dao.Prepare(meetingInsertStmt)
	entity.CheckErr(err)	
	if err != nil {
        return err
    }
    defer stmt.Close()
	_, err2 := stmt.Exec(meeting.Title, meeting.Sponsor, meeting.Participants.String, meeting.StartTime, meeting.EndTime);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//delete a meeting by title
var meetingDeleteByTimeStmt = "delete from Meeting where Title = ?"
func (dao *MeetingInfoDao) DeleteAMeetingByTitle(title string ) error {
    stmt, err := dao.Prepare(meetingDeleteByTimeStmt)
	entity.CheckErr(err)	
	if err != nil {
        return err
    }
    defer stmt.Close()
	_, err2 := stmt.Exec(title);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//delete all sponsor meeting
var deleteAllSponsorMeetingStmt = "delete from Meeting where Sponsor = ?"
func (dao *MeetingInfoDao) DeleteAllSponsorMeeting(name string ) error {
    stmt, err := dao.Prepare(deleteAllSponsorMeetingStmt)
	entity.CheckErr(err)	
	if err != nil {
        return err
    }
    defer stmt.Close()
	_, err2 := stmt.Exec(name);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}
//update a meeting's participants
var meetingParticipantsUpdateStmt = "UPDATE Meeting SET Participants = ? WHERE Title = ? "
func (dao *MeetingInfoDao) UpdateMeetingParticipants(title, Participants string) error {
    stmt, err := dao.Prepare(meetingParticipantsUpdateStmt)
	entity.CheckErr(err)	
	if err != nil {
        return err
    }
    defer stmt.Close()
	_, err2 := stmt.Exec(Participants, title);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//get all participant meeting
func (dao *MeetingInfoDao) GetAllParticipantsMeeting(name string) ([]entity.Meeting, error)  {
	stmt := "select * from Meeting where Participants like '%" + name + "%'"
	rows, _ := dao.Query(stmt)
	defer rows.Close()
	var meetingSlice = make([]entity.Meeting, 0)
	for rows.Next() {
		var meeting = entity.Meeting{}
		err := rows.Scan(&meeting.Id, &meeting.Title, &meeting.Sponsor, &meeting.Participants, &meeting.StartTime, &meeting.EndTime)
        meetingSlice = append(meetingSlice, meeting)
        entity.CheckErr(err)
        if err != nil {
            return nil, err
        }
	}
    return meetingSlice, nil
}


//get all relative meeting
var getAllRelativeMeetingStmt = "select * from Meeting where Sponsor = ? or Participants like '%zhangzemian%'"
func (dao *MeetingInfoDao) GetAllRelativeMeeting(name string) ([]entity.Meeting, error)  {
	stmt := "select * from Meeting where Sponsor = \"" + name + "\" or Participants like '%" + name + "%'"
	rows, _ := dao.Query(stmt)
	defer rows.Close()
	var meetingSlice = make([]entity.Meeting, 0)
	for rows.Next() {
		var meeting = entity.Meeting{}
		err := rows.Scan(&meeting.Id, &meeting.Title, &meeting.Sponsor, &meeting.Participants, &meeting.StartTime, &meeting.EndTime)
        meetingSlice = append(meetingSlice, meeting)
        entity.CheckErr(err)
        if err != nil {
            return nil, err
        }
	}
    return meetingSlice, nil
}

//Get AllMeeting Between StartTime And EndTime Of someone
func (dao *MeetingInfoDao) GetAllMeetingBetweenStartTimeAndEndTimeOfSomeone(name string, startTime time.Time, endTime time.Time) ([]entity.Meeting, error)  {
	allRelativeMeeting, err := dao.GetAllRelativeMeeting(name)
	entity.CheckErr(err)
	if err != nil {
		return nil, err
	}
	var rightmeetingSlice = make([]entity.Meeting, 0)
	for _, meeting := range allRelativeMeeting {
		if small_date(meeting.EndTime, startTime) || large_date(meeting.StartTime, endTime) {
			;
		} else {
			rightmeetingSlice = append(rightmeetingSlice, meeting)
		}
	}
    return rightmeetingSlice, nil
}

//find meeting by id
var getMeetingByID = "select * from Meeting where Id = ?"
func (dao *MeetingInfoDao) GetMeetingByID(id string) (*entity.Meeting, error)  {
	fmt.Println(id)
	var meeting = entity.Meeting{}
    stmt, err := dao.Prepare(getMeetingByID)
	entity.CheckErr(err)
	if err != nil {
        return nil, err
    }
	defer stmt.Close()
	result := stmt.QueryRow(id)
	err2 := result.Scan(&meeting.Id, &meeting.Title, &meeting.Sponsor, &meeting.Participants, &meeting.StartTime, &meeting.EndTime)

	entity.CheckErr(err2)
	if err2 != nil {
		return nil, err2
	}
	return &meeting, nil
}

//find meeting by Title
var getMeetingByTitle = "select * from Meeting where Title = ?"
func (dao *MeetingInfoDao) GetMeetingByTitle(title string) (*entity.Meeting, error)  {
	var meeting = entity.Meeting{}
    stmt, err := dao.Prepare(getMeetingByTitle)
	entity.CheckErr(err)
	if err != nil {
        return nil, err
    }
	defer stmt.Close()
	result := stmt.QueryRow(title)
	err2 := result.Scan(&meeting.Id, &meeting.Title, &meeting.Sponsor, &meeting.Participants, &meeting.StartTime, &meeting.EndTime)

	entity.CheckErr(err2)
	if err2 != nil {
		return nil, err2
	}
	return &meeting, nil
}
//time process method
func small_date(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}

func large_date(date1, date2 time.Time) bool {
	return date1.After(date2) || date1.Equal(date2)
}

func Date_to_string(date time.Time) string {
	return date.Format("2006-01-02/15:04")
}

func String_to_date(date string) (time.Time, error) {
	the_time, err := time.Parse("2006-01-02/15:04", date)
	return the_time, err
}