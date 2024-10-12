package db

import (
	"m3u82mp4/consts/errcode"

	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	Id         int64  `gorm:"column:id;primaryKey"`
	Account    string `gorm:"column:account"`
	Password   string `gorm:"column:password"`
	Username   string `gorm:"column:username"`
	Avatar     string `gorm:"column:avatar"`
	IsDel      bool   `gorm:"column:is_del"`
	Createtime int64  `gorm:"column:createtime;autoCreateTime:milli"`
	Updatetime int64  `gorm:"column:updatetime;autoUpdateTime:milli"`
}

type UserDB struct {
	db *gorm.DB
}

func table() string {
	return "v_user"
}

func NewUserDB(db *gorm.DB) *UserDB {
	userdb := new(UserDB)
	userdb.db = db
	return userdb
}

func (u *UserDB) GetAccountCount(account string) int64 {
	var count int64
	u.db.Table(table()).Where("account = ?", account).Count(&count)
	return count
}

func (u *UserDB) GetUserInfoByAccount(account string) *User {
	var user User
	u.db.Table(table()).Where("account = ?", account).First(&user)
	return &user
}

func (u *UserDB) GetUserInfoById(id int64) *User {
	var user User
	u.db.Table(table()).Where("id = ?", id).First(&user)
	return &user
}

func (u *UserDB) Save(user *User) *errcode.ErrorCode {
	rt := u.db.Table(table()).Save(user)
	if rt.Error != nil {
		return &errcode.USER_SAVE_ERR
	}
	return nil
}
