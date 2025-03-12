package db

import (
	"m3u82mp4/consts/errcode"

	"gorm.io/gorm"
)

// Message 留言板
type Message struct {
	Id         int64      `gorm:"column:id;primaryKey" json:"id"`
	Content    string     `gorm:"column:content" json:"content"`
	ParentId   int64      `gorm:"column:parent_id"  json:"parentId"`
	IsDel      bool       `gorm:"column:is_del"  json:"isDel"`
	UserID     int64      `gorm:"column:user_id"  json:"userId"`
	Username   string     `gorm:"column:username"  json:"username"`
	Avatar     string     `gorm:"column:avatar"  json:"avatar"`
	Createtime int64      `gorm:"column:createtime;autoCreateTime:milli" json:"createTime"`
	Updatetime int64      `gorm:"column:updatetime;autoUpdateTime:milli" json:"updateTime"`
	Replies    []*Message `gorm:"-" json:"replies"` // 存储回复列表，不映射到数据库
}

type MessageDB struct {
	db *gorm.DB
}

func (Message) TableName() string {
	return "v_message"
}

func NewMessageDB(db *gorm.DB) *MessageDB {
	messagedb := new(MessageDB)
	messagedb.db = db
	return messagedb
}

// Create 创建留言
func (m *MessageDB) Create(message *Message) *errcode.ErrorCode {
	err := m.db.Create(message).Error
	if err != nil {
		errCustom := errcode.DBCustom(err.Error())
		return &errCustom
	}
	return nil
}

// Delete 删除留言
func (m *MessageDB) Delete(id int64, userId int64) *errcode.ErrorCode {
	err := m.db.Model(&Message{}).Where("id = ? AND user_id = ?", id, userId).
		Update("is_del", true).Error
	if err != nil {
		errCustom := errcode.DBCustom(err.Error())
		return &errCustom
	}
	return nil
}

// QueryList 查询留言列表
func (m *MessageDB) QueryList() (*errcode.ErrorCode, []*Message) {
	var messages []*Message
	err := m.db.Where("is_del = ?", false).
		Order("createtime desc").
		Find(&messages).Error
	if err != nil {
		errCustom := errcode.DBCustom(err.Error())
		return &errCustom, nil
	}
	return nil, messages
}

// QueryById 根据ID查询留言
func (m *MessageDB) QueryById(id int64) (*errcode.ErrorCode, *Message) {
	var message Message
	err := m.db.Where("id = ? AND is_del = ?", id, false).First(&message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		errCustom := errcode.DBCustom(err.Error())
		return &errCustom, nil
	}
	return nil, &message
}
