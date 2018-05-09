package utils


func StartIndex(page,pageSize int)(pageIndex int){
	if page < 1 {
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
