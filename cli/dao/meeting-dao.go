package dao
import (
//		"errors"
		"time"
		"entity"
	)

type MeetingInfoDao entity.DaoSource

var meetingInsertStmt = "Insert into Meeting values (?,?,?,?,?,?)"
//InsertUser
func (dao *MeetingInfoDao) InsertMeeting(username string, meeting entity.Meeting) error {
    stmt, err := dao.Prepare(meetingInsertStmt)
	entity.CheckErr(err)
	if err != nil {
        return err
    }
    defer stmt.Close()

	_, err2 := stmt.Exec(meeting.Id, meeting.Title, meeting.Sponsor, meeting.Participants, meeting.StartTime, meeting.EndTime);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}


var getMeetingBetweenDateStmt = "select from Meeting where Id = ? and StartTime >= ? and EndTime <= ?"
//InsertUser
func (dao *MeetingInfoDao) GetAllMeetingBetweenStartTimeAndEndTimeOfAnyUser(idSlice []string, startTime time.Time, endTime time.Time) ([]entity.Meeting, error)  {
    stmt, err := dao.Prepare(meetingInsertStmt)
	entity.CheckErr(err)
	if err != nil {
        return nil, err
    }
	defer stmt.Close()
	var meetingSlice = make([]entity.Meeting, 0)
	// 有错误
	// for index, value := range idSlice {
	// 	_, err2 := stmt.QueryRow(value, startTime, endTime)
	// 	entity.CheckErr(err2)
	// 	if err2 != nil {
	// 		return nil, err2
	// 	}
	// 	meetingSlice = append(meetingSlice, )
	// }
	// 将就版
	for _, value := range idSlice {
		err2 := stmt.QueryRow(value, startTime, endTime)
		//entity.CheckErr(err2)
		if err2 != nil {
			//return nil, err2
		}
		meetingSlice = append(meetingSlice, )
	}
    return nil, nil
}

//time process method
func small_date(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}

func large_date(date1, date2 time.Time) bool {
	return date1.After(date2) || date1.Equal(date2)
}