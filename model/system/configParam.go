package system

type ConfigParam struct {
	MT      string `json:"mt"`      // 方法类型
	Id      int64  `json:"id"`      // ID
	Key     string `json:"key"`     // 键名
	Value   string `json:"value"`   // 键值
	Type    string `json:"type"`    // 类型
	State   uint8  `json:"state"`   // 状态 1启用 0禁用
	Comment string `json:"comment"` // 详情说明
}
