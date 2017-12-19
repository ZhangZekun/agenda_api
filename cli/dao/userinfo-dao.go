package dao
import (
        "errors"
        "agenda_api/cli/entity"
	)

type UserInfoDao entity.DaoSource

var userInfoQueryByUsernameStmt = "SELECT * FROM User where Username = ?"
// FindByUsername
func (dao *UserInfoDao) FindByUsername(name string) (*entity.User, error) {
    stmt, err := dao.Prepare(userInfoQueryByUsernameStmt)
    entity.CheckErr(err)
    defer stmt.Close()

    row := stmt.QueryRow(name)
	u := entity.User{}
    err = row.Scan(&u.Username, &u.Password, &u.Email, &u.Phone, &u.SponsorMeeting, &u.ParticipantMeeting);
	if err != nil {
		return nil, err
	}
    return &u, nil
}

var userInfoInsertStmt = "Insert into User values (?,?,?,?,?,?)"
//InsertUser
func (dao *UserInfoDao) InsertUser(user entity.User) error {
	if user,_ := dao.FindByUsername(user.Username); user != nil {
		err := errors.New("Such Username have been registered! Please try another username");
		return err
	}
    stmt, err := dao.Prepare(userInfoInsertStmt)
	entity.CheckErr(err)
	if err != nil {
        return err
    }
    defer stmt.Close()

	_, err2 := stmt.Exec(user.Username, user.Password, user.Email.String, user.Phone.String, user.SponsorMeeting.String, user.ParticipantMeeting.String);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//determine whether a user has been logged in by username or not
var UserHasLoginStmt = "SELECT * FROM LoginUsers where Username = ?"
func (dao *UserInfoDao) UserHasLogin(name string) bool {
    stmt, err := dao.Prepare(UserHasLoginStmt)
    entity.CheckErr(err)
    defer stmt.Close()

    rows, err2 := stmt.Query(name)
    entity.CheckErr(err2)
    defer rows.Close()
    if rows.Next() {     
        return true;
    }
    return false;
}

//Insert User' login infomation
var userLoginInfoInsertStmt = "Insert into LoginUsers values (?)"
func (dao *UserInfoDao) LoginInfoInsert(username string) error {
	if user,_ := dao.FindByUsername(username); user == nil {
		err := errors.New("No such username, login fail");
		return err
    }
    if dao.UserHasLogin(username) {
        err := errors.New("You have logged in already, no need to log in again");
		return err
    }
    stmt, err := dao.Prepare(userLoginInfoInsertStmt)
	entity.CheckErr(err)
	if err != nil {
        return err
    }
    defer stmt.Close()

	_, err2 := stmt.Exec(username);
	
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//delete User's log in infomation by username
var userLoginInfoDeleteStmt = "Delete from LoginUsers where Username=?"
func (dao *UserInfoDao) LoginInfoDelete (username string) error {
	if user,_ := dao.FindByUsername(username); user == nil {
		err := errors.New("no such user in record!");
		return err
    }
    if !dao.UserHasLogin(username) {
        err := errors.New("You haven't logged in, no need to log out");
		return err
    }
    stmt, err := dao.Prepare(userLoginInfoDeleteStmt)
	entity.CheckErr(err)
	if err != nil {
        return err
    }
    defer stmt.Close()

	_, err2 := stmt.Exec(username);
    entity.CheckErr(err2)
    if err2 != nil {
        return err2
    }
    return nil
}

//get all users' infomation
var getAllUsersInfoStmt = "SELECT * FROM User where Username != ?"
// FindByUsername
func (dao *UserInfoDao) GetAllUsersInfo(username string) ([]entity.User, error) {
    if login := dao.UserHasLogin(username); login == false {
        err := errors.New("you haven't log in! please log in first!")
        return nil, err
    }
    stmt, err := dao.Prepare(userInfoQueryByUsernameStmt)
    entity.CheckErr(err)
    defer stmt.Close()

    rows, err := stmt.Query("")
    defer rows.Close()
    entity.CheckErr(err)
    if err != nil {
		return nil, err
    }
    
    userSlice := make([]entity.User, 0)
    for rows.Next() {
        var u = entity.User{}
        err = rows.Scan(&u.Username, &u.Password, &(u.Email), &u.Phone, &u.SponsorMeeting, &u.ParticipantMeeting);
        userSlice = append(userSlice, u)
        entity.CheckErr(err)
        if err != nil {
            return nil, err
        }
    }
    return userSlice, nil
}