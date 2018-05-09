package models

//响应体
type  BaseMsgResp struct {
	Ret int //状态
	Msg string //信息
	Object interface{}
}

