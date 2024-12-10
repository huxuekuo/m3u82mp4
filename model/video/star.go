package video

type StarParam struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
}

func (s *StarParam) Check() bool {
	if s.ID == "" || len(s.ID) <= 0 {
		return false
	}
	return true
}
