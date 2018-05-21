package utils

import (
	"mime/multipart"
	"strconv"
	"io"
	"os"
	"strings"
	"fmt"
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