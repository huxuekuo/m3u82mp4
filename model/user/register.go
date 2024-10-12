package user

type RegisterParam struct {
	Account  string `json:"u" form:"u"`
	PassWord string `json:"p" form:"p"`
}

// 检查参数
func (r *RegisterParam) Check() (bool, string) {
	if len(r.Account) <= 0 {
		return true, "账号为空"
	}
	if len(r.PassWord) <= 0 {
		return true, "密码为空"
	}
	return false, ""
}

type LoginParam struct {
	RegisterParam
}
