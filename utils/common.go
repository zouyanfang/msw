package utils

import (
	"mime/multipart"
	"strconv"
	"io"
	"os"
	"strings"
	"fmt"
	"math/rand"
	"time"
)

func StartIndex(page,pageSize int)(pageIndex int){
	if page <= 1 {
		pageIndex = 0
		return
	}
	pageIndex = (page-1)*pageSize
	return
}

func PageCount(pageSize,count int)(pages int){
	if count%pageSize==0 {
		pages = count/pageSize
		return
	}
	pages = count /pageSize+1
	return
}

func IsHaveNext(page,pages int)(bool){
	if page >= pages {
		return false
	}
	return true
}

//图片保存
func SaveImg(s multipart.File,f *multipart.FileHeader,id int,path string)(name string){
	mode := strings.Split(f.Filename,".")
	newsname := strconv.Itoa(id)+"."+mode[len(mode)-1]
	os.Remove("static/img/"+path+newsname)
	f1,err := os.Create("static/img/"+path+newsname)
	if err != nil {
		return
	}
	_,err = io.Copy(f1,s)
	if err != nil {
		fmt.Println(err.Error())
	}
	f1.Close()
	name = newsname
	return name
}

//获取图片的数字
func GetRegisterImgPosition(img []string)(right []string){
	length := len(img)
	a := make([]string,0)
	number := make([]int,0)
	types := make([]string,0)
	for i:=0;i<length;i++ {
		str := strings.Split(img[i],"/")
		number := strings.Split(str[len(str)-1],".")
		fmt.Println("图片数字",number[0])
		a = append(a,number[0])
		types = append(types,number[len(number)-1])
	}
	fmt.Println("number",number)
	fmt.Println("type","type")
	length = len(a)
	for i := 0;i < length;i++ {
		s,_ := strconv.Atoi(a[i])
		number = append(number,s)
	}
	fmt.Println("len(number)",len(number))
	SortNumber(number)
	fmt.Println("排序过的图片",number)
	for i := 0;i<length;i++{
		str := "../static/img/step/"+strconv.Itoa(number[i])+"."+types[i]
		fmt.Println("还原图片的字符串",str)
		right = append(right,str)
	}
	return
}

func SortNumber(a []int)([]int){
	length := len(a)
	fmt.Println("number长度",length)
	for i := 0; i<length; i++{
		for j := 0; j < length-1-i;j++ {
			if a[j] > a[j+1]{
				a[j+1],a[j] = a[j],a[j+1]
			}
		}
	}
	fmt.Println("排序后的图片",a)
	return a
}

//随机默认图片
func RandUserImg()(int){
	 rand.Seed(time.Now().Unix())
	 n := rand.Intn(5)
	 return n
}