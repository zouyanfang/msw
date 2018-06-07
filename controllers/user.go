package controllers

import (
	"msw/models"
	"msw/services"
	"msw/utils"
	"time"
	"fmt"
	"strings"
	"sync"
)

type UserController struct {
	BaseController
}

//用户菜谱页面


func (this *UserController)UserCenter(){
	this.Data["type"] = 1
	this.IsNeedTemplate()
	resp := services.GetUserDish(this.User.Id)
	fmt.Println(resp.CollectDish)
	fmt.Println()
	this.Data["resp"] = resp
	this.TplName = "site/user/dishcenter.html"
}

//用户菜单页面
func (this *UserController)UserMenu(){
	this.Data["type"] = 2
	this.IsNeedTemplate()
	resp := services.GetMenuListByUid(this.User.Id)
	fmt.Println(resp)
	this.Data["menu"] = resp
	this.TplName = "site/user/menucenter.html"
}

//用户个人信息设置页面
func (this *UserController)UserConfig(){
	this.Data["type"] = 3
	this.IsNeedTemplate()
	this.TplName = "site/user/userconfig.html"
}


//用户模板
func (this *UserController)IsNeedTemplate(){
	this.IsHaveUser()
	this.Data["user"] = this.User
	this.Data["dishcount"] = models.GetDishCountByUid(this.User.Id)
	this.Data["menucount"] = models.GetMenuCountByUid(this.User.Id)
	this.Layout = "site/user/template.html"
}

//判断是否有用户
func (this *UserController)IsHaveUser(){
	if this.User == nil {
		this.Redirect("/login",302)
		return
	}
}

//test
func (this *UserController)Demo(){
	this.TplName = "site/user/1.html"
}

//test
func (this *UserController)GetImg(){
	//_,s,_ := this.GetFile("file")
	this.ServeJSON()
}

//发布评论
func (this *UserController)ReleaseTalk(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if (this.User == nil ){
		resp.Msg = "用户还未登入，请先登入"
		return
	}
	content := this.GetString("content")
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = "获取参数失败"
		return
	}
	resp = services.ReleaseDishTalk(dishid,this.User.Id,content)
	return
}

//用户收藏菜谱
func (this *UserController)UserCollection(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登录！"
		return
	}
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	status,err := this.GetInt("status")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp  = services.UserDishCollect(dishid,this.User.Id,status)
}

//获取用户的菜单
func (this *UserController)GetUserMenu(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
		return
	}
	resp = services.GetUserMenu(this.User.Id)
	return
}

//添加至菜单
func (this *UserController)AddToMenu(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登入！"
		return
	}
	dishid,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	menuid,err := this.GetInt("menuid")
	resp = services.AddToMenu(menuid,dishid)
}


//创建菜单
func (this *UserController)CreateNewImg(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
	}
	s,f,err := this.GetFile("img")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	id := int(time.Now().Unix())
	path := utils.SaveImg(s,f,id,"menu_img/")
	path = "../static/img/menu_img/"+path
	menuname := this.GetString("menuname")
	describe := this.GetString("describe")
	resp = services.CreateNewMenu(this.User.Id,menuname,path,describe)
}

//发布留言板留言
func (this *UserController)ReleaseUserMsg(){
	var resp models.BaseMsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
		return
	}
	content:= this.GetString("content")
	if content == ""{
		resp.Msg = "评论内容为空！"
		return
	}
	resp = services.ReleaseUserMsg(this.User.Id,content)
}

//用户收藏菜单
func (this *UserController)UserCollectionMenu(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func(){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登录"
		return
	}
	menuid,err := this.GetInt("menuid")
	if err != nil {
		resp.Msg = err.Error()+"adsasdasdsad"
		return
	}
	status,err := this.GetInt("status")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp = services.UserMenuCollect(menuid,this.User.Id,status)
	fmt.Println(resp)
	return
}

//删除菜单
func (this *UserController)DeleteMenu()(err error){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
		return
	}
	menuid,err := this.GetInt("menuid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp = services.DeleteMenu(menuid,this.User.Id)
	return
}

//修改用户基础信息
func (this *UserController)ModifyBaseMsg(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登入"
		return
	}
	condition := ""
	paras := []interface{}{}
	if name:=this.GetString("name"); name != ""{
		condition = " SET name = ? "
		paras = append(paras,name)
	}
	if status,_ :=this.GetInt("status");status == 1{
		s,f,err := this.GetFile("file")
		if err != nil {
			resp.Msg = err.Error()
			return
		}
		id := int(time.Now().Unix())
		path := utils.SaveImg(s,f,id,"user_img")
		path = "../static/img/user_img"+path
		condition += " SET user_img = ? "
		paras = append(paras,path)
	}
	oldpwd := this.GetString("old")
	if oldpwd != ""{
		if oldpwd != this.User.Pwd{
			resp.Msg = "原始密码错误 Orz"
			return
		}
		pwd := this.GetString("pwd")
		if pwd != ""{
			condition += " SET pwd = ? "
			paras = append(paras,pwd)
		}
	}
	resp = services.ModifyUserByUid(this.User.Id,condition,paras)
}





func (this *UserController)CreateDish(){
	var resp models.CreateDishResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
	}
	s,f,err := this.GetFile("img")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	id := int(time.Now().Unix())
	path := utils.SaveImg(s,f,id,"dish_img/")
	path = "../static/img/dish_img/"+path
	dishname := this.GetString("dishname")
	describe := this.GetString("describe")
	resp = services.CreateNewDish(this.User.Id,dishname,path,describe)
	fmt.Println(resp)
}

func (this *UserController)UpdateDish(){
	var resp models.BaseMsgResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	main := this.GetString("main")
	second := this.GetString("second")
	taste := this.GetString("taste")
	system := this.GetString("system")
	dish,_:=this.GetInt("dishid")
	resp = services.UpdateDishDetail(dish,taste,system,main,second)
}


func (this *UserController)UpdateDishStep(){
	var resp models.BaseMsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	step,_ := this.GetInt("step")
	desc := this.GetString("describe")
	dishid,_ := this.GetInt("dishid")
	str := strings.Split(desc,"/")
	fmt.Println(str)
	j := step-1
	fmt.Println("j",j)
	for i:=step;i>= 1;i--{
		if  j<0 {
			break;
		}
		err := models.InsertDishStep(dishid,i,str[j])
		fmt.Println("i:",i,"str",str[j])
		if err != nil {
			resp.Msg = err.Error()
			return
		}
		j--
	}
	resp.Ret = 200
}

var l sync.RWMutex

var Step int

func (this *UserController)GetStepImg(){
	s,f,err := this.GetFile("file")
	if err != nil {
		return
	}
	//该方法只能被一个线程访问
	l.Lock()
	id := int(time.Now().Unix())
	time.Sleep(1*time.Second)
	fmt.Println("过一秒")
	path := utils.SaveImg(s,f,id,"step/")
	path = "../static/img/step/"+path
	ids,_ := models.GetLastDish()
	Step,_ = models.CountStep(ids)
	models.InsertImgstep(ids,path,Step)
	imgs,_ := models.GetImg(ids)
	fmt.Println("存入的图片长度",len(imgs))
	rightposition := utils.GetRegisterImgPosition(imgs)
	length := len(rightposition)
	count := models.GetCountImg(ids)
	fmt.Println("count",count,"length",length)
	if count == length {
		j := 0
		for i :=1 ;i <= count ;i++{
			models.UpdateRight(ids,rightposition[j],i)
			j++
		}

	}
	fmt.Println("2331231231231231231231231231")
	l.Unlock()
	this.Data["ok"] = "ok"
	this.ServeJSON()
}


func (this *UserController)Delete(){
	var resp models.BaseMsgResp
	defer func(){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,_ := this.GetInt("id")
	err := models.DeleteDish(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
}


func (this *UserController)CancelDish(){
	var resp models.BaseMsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,_ := this.GetInt("dishid")
	resp = services.UserDishCollect(id,this.User.Id,0)
}