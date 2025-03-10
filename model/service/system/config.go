package systemser

import (
	"m3u82mp4/consts/errcode"
	"m3u82mp4/model/common"
	"m3u82mp4/model/db"
	"m3u82mp4/model/system"
	"m3u82mp4/utils"

	"gorm.io/gorm"
)

type ConfigService struct {
	db *gorm.DB
}

func NewConfigService(db *gorm.DB) *ConfigService {
	return &ConfigService{
		db: db,
	}
}

const (
	CONFIG_TYPE_DEFAULT = ""
)

const (
	CONFIG_KEY_DEFAULT = ""
)

func (c ConfigService) APIAdd(param *system.ConfigParam) any {
	err, id := utils.ID()
	res := &common.Respone{}
	if err != nil {
		return res.Set(errcode.SYS_RETRY)
	}
	item := &db.SystemConfig{
		Type:  param.Type,
		Key:   param.Key,
		Value: param.Value,
		Id:    id,
	}
	configdb := db.NewSystemConfigDB(c.db)
	errcode, b := configdb.Create(item)
	if errcode != nil {
		return res.Set(*errcode)
	}
	return res.OK(b)
}

func (c ConfigService) APIQueryList(param *system.ConfigParam) any {
	res := &common.Respone{}
	res.OK(c.QueryList(param))
	return res
}

func (c ConfigService) QueryList(param *system.ConfigParam) []*db.SystemConfig {
	configdb := db.NewSystemConfigDB(c.db)
	errcode, list := configdb.QueryList(param.Type, param.Key)
	if errcode != nil {
		return nil
	}
	return list
}
