package video

type PlayRecordParam struct {
	StartTime string `json:"startTime" form:"startTime"`
	Index     string `json:"index" form:"index"`
	Teleplay  string `json:"teleplay" form:"teleplay"`
	Name      string `json:"name" form:"name"`
}
