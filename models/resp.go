package models

//响应体
type BaseMsgResp struct {
	Ret    int    //状态
	Msg    string //信息
	Object interface{}
}

//菜谱返回体
type DishResp struct {
	BaseMsgResp
	Page int
	StepInfo
	Count int
	Next bool
	Pref bool

}

//菜谱详情响应体
type DishDetailResp struct {
	BaseMsgResp
	DishDetail     *DishInfo
	StepDetail     []StepInfo
	MainMaterial   []string
	SecondMaterial []string
}

type UserCollectResp struct {
	BaseMsgResp
	Status int
}

type MenuResp struct {
	BaseMsgResp
	Page int
}

type MenuDetailResp struct {
	BaseMsgResp
	Name string
	Img  string
	Dish []Dish
}

type SearchResp struct {
	BaseMsgResp
	Menu []Menu
	Dish []Dish
	Status bool
}

type UserMsgResp struct {
	BaseMsgResp
	Page int
	Count int
}

type UserInfoIndexResp struct {
	BaseMsgResp
	Dish []Dish
}
