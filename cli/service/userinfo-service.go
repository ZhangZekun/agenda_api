package service

import (
    "agenda_api/cli/entity"
    "agenda_api/cli/dao"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}


// FindByUsername .
func (*UserInfoAtomicService) FindByUsername(name string) (*entity.User, error) {
    dao := dao.UserInfoDao{entity.Mydb}
    return dao.FindByUsername(name)
}

//InsertAUser
func (*UserInfoAtomicService) InsertUser(user entity.User) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.InsertUser(user)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//Insert Login Info 
func (*UserInfoAtomicService) LoginInfoInsert(username string) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.LoginInfoInsert(username)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//delete User's log in infomation by username
func (*UserInfoAtomicService) LoginInfoDelete(username string) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.LoginInfoDelete(username)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}
//determine whether a user has been logged in by username or not
func (*UserInfoAtomicService) UserHasLogin(name string) bool {
    dao := dao.UserInfoDao{entity.Mydb}
    return dao.UserHasLogin(name)
}

//get all users' infomation
func (*UserInfoAtomicService) GetAllUsersInfo(name string) ([]entity.User, error) {
    dao := dao.UserInfoDao{entity.Mydb}
    return dao.GetAllUsersInfo(name)
}
