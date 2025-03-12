package videoser

import (
	"gorm.io/gorm"
)

type VideoSer struct {
	db *gorm.DB
}

func NewVideoSer(db *gorm.DB) *VideoSer {
	instance := new(VideoSer)
	instance.db = db
	return instance
}

func (vs *VideoSer) GetMonitorList() {
	// monitorDB := db.NewMonitorDB(vs.db)
	// monitorList, errc := monitorDB.List()
	// if errc != nil {
	// }

	// for _, v := range monitorList {
	// }
}
