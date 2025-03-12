package message

import (
	"m3u82mp4/consts/errcode"
	"m3u82mp4/model/common"
	"m3u82mp4/model/db"
	"m3u82mp4/utils"

	"gorm.io/gorm"
)

type MessageService struct {
	db *gorm.DB
}

func NewMessageService(db *gorm.DB) *MessageService {
	return &MessageService{db: db}
}

// Add 添加留言
func (s *MessageService) Add(content string, parentId int64, userId int64, username string, avatar string) any {
	res := &common.Respone{}

	// 如果是回复，检查父留言是否存在
	if parentId > 0 {
		messageDB := db.NewMessageDB(s.db)
		code, parentMessage := messageDB.QueryById(parentId)
		if code != nil {
			return res.Set(*code)
		}
		if parentMessage == nil {
			return res.Set(errcode.ApiCustom("要回复的留言不存在"))
		}
	}

	err, id := utils.ID()
	if err != nil {
		return res.Set(errcode.ApiCustom("生成ID失败"))
	}

	message := &db.Message{
		Id:       id,
		Content:  content,
		ParentId: parentId,
		UserID:   userId,
		Username: username,
		Avatar:   avatar,
	}

	messageDB := db.NewMessageDB(s.db)
	if err := messageDB.Create(message); err != nil {
		return res.Error("创建留言失败")
	}
	return res.OK(nil)
}

// Delete 删除留言
func (s *MessageService) Delete(id int64, userId int64) any {
	res := &common.Respone{}
	messageDB := db.NewMessageDB(s.db)
	if err := messageDB.Delete(id, userId); err != nil {
		return res.Error("删除留言失败")
	}
	return res.OK(nil)
}

// QueryList 查询留言列表
func (s *MessageService) QueryList() any {
	res := &common.Respone{}
	messageDB := db.NewMessageDB(s.db)
	code, messages := messageDB.QueryList()
	if code != nil {
		return res.Error("查询留言列表失败")
	}

	// 构建留言树
	messageMap := make(map[int64]*db.Message)
	var rootMessages []*db.Message

	for _, msg := range messages {
		messageMap[msg.Id] = msg
		if msg.ParentId == 0 {
			rootMessages = append(rootMessages, msg)
		}
	}

	// 组织回复
	for _, msg := range messages {
		if msg.ParentId > 0 {
			if parent, ok := messageMap[msg.ParentId]; ok {
				if parent.Replies == nil {
					parent.Replies = make([]*db.Message, 0)
				}
				parent.Replies = append(parent.Replies, msg)
			}
		}
	}

	return res.OK(rootMessages)
}
