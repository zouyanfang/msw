
	$("#foodmenu").click(function(){
		if ($(this).hasClass("glyphicon-plus")){
			$(this).removeClass("glyphicon-plus")
			$(this).addClass("glyphicon-minus")
			return
		}		
		$(this).addClass("glyphicon-plus")
		$(this).removeClass("glyphicon-minus")
	})

	$("#taste .menuitems").click(function(){
		// 移动点全部的背景图
		$("#taste .menuitems").css("background-color","white")
		$("#taste .menuitems").removeClass("white")
		var s = $(this).css("color")
		$(this).css("background-color",s)
		$(this).addClass("white")
		var text = $(this).text()
		var p = `<p class="menuitems col-sm-2 text-center taste" ><span class="glyphicon glyphicon-remove todelete"></span>`
		p +=text+"</p>"
		$("#condition").find(".taste").remove()
		$("#condition").append(p)
			var area = ""
		if ($("#condition .area")){
			area = $("#condition .area").text()
		}
		var taste = ""
		if ($("condition .taste")){
			taste = $("#condition .taste").text()
		}
		$.ajax({
			url: '/dish/getconditiondishlist',
			type: 'get',
			dataType: 'json',
			data: {"tasty":taste,"area":area},
			success:function(result){
				console.log(result)
				if (result.Ret == 200){
					$(".dishcontent").empty()
					if (result.Object != null){
						for(var i = 0;i<result.Object.length;i++){
							var text = "<a href='/dish/getdishdetail?uid="+result.Object[i].Uid+"&dishId="+result.Object[i].Id+"' class='col-sm-4 sgo' id='food'>"+
											"<div class='col-sm-12' style='padding: 0px;'><img class='img-responsive' src='"+result.Object[i].DishImg+"'></div>"+
												"<div class='col-sm-12' >"+
											"<div class='col-sm-4'>"+
											"<img src='"+result.Object[i].UserImg+"'  height='50px;' width='50px;' style='border-radius: 50px;'>"+
											"</div>"+
											"<div class='col-sm-8'>"+
											"<i class='col-sm-12'>"+result.Object[i].Name+"</i>"+
											"<i class='col-sm-12'>"+result.Object[i].DishName+"</i>"+
											"</div>"+
											"</div>"+
										"</a>"
							$(".dishcontent").append(text)
						}
					}else {
						$(".dishcontent").append(`<div class="alert alert-danger col-sm-12" role="alert">
													<strong>o(╯□╰)o</strong> 没有该搜索条件的菜谱
												</div>`)
						}
					
					if (result.Pref){
						$("#prev").removeClass('disabled')
					}else{	
						console.log(result.Pref)
						$("#prev").addClass('disabled')
					}
					if(result.Next){
						$("#next").removeClass('disabled')
					}else{
						$("#next").addClass('disabled')
					}
					$("#page").text(result.Page)
				}
			}
		})

	})

	$("#area .menuitems").click(function(){
		$("#area .menuitems").css("background-color","white")
		$("#area .menuitems").removeClass("white")
		var s = $(this).css("color")
		$(this).css("background-color",s)
		$(this).addClass("white")
		var text = $(this).text()
		var p = `<p class="menuitems col-sm-2 text-center area" ><span class="glyphicon glyphicon-remove todelete"></span>`
		p +=text+"</p>"
		$("#condition").find(".area").remove()
		$("#condition").append(p)
		var area = ""
		if ($("#condition .area")){
			area = $("#condition .area").text()
		}
		var taste = ""
		if ($("condition .taste")){
			taste = $("#condition .taste").text()
		}
		$.ajax({
			url: '/dish/getconditiondishlist',
			type: 'get',
			dataType: 'json',
			data: {"tasty":taste,"area":area},
			success:function(result){
				console.log(result)
				if (result.Ret == 200){
					$(".dishcontent").empty()
					if (result.Object != null){
					for(var i = 0;i<result.Object.length;i++){
						var text = "<a href='/dish/getdishdetail?uid="+result.Object[i].Uid+"&dishId="+result.Object[i].Id+"' class='col-sm-4 sgo' id='food'>"+
										"<div class='col-sm-12' style='padding: 0px;'><img class='img-responsive' src='"+result.Object[i].DishImg+"'></div>"+
											"<div class='col-sm-12' >"+
										"<div class='col-sm-4'>"+
										"<img src='"+result.Object[i].UserImg+"'  height='50px;' width='50px;' style='border-radius: 50px;'>"+
										"</div>"+
										"<div class='col-sm-8'>"+
										"<i class='col-sm-12'>"+result.Object[i].Name+"</i>"+
										"<i class='col-sm-12'>"+result.Object[i].DishName+"</i>"+
										"</div>"+
										"</div>"+
									"</a>"
						$(".dishcontent").append(text)
					}
				}else {
						$(".dishcontent").append(`<div class="alert alert-danger col-sm-12" role="alert">
													<strong>o(╯□╰)o</strong> 没有该搜索条件的菜谱
												</div>`)
						}
					if (result.Pref){
						$("#prev").removeClass('disabled')
					}else{	
						console.log(result.Pref)
						$("#prev").addClass('disabled')
					}
					if(result.Next){
						$("#next").removeClass('disabled')
					}else{
						$("#next").addClass('disabled')
					}
					$("#page").text(result.Page)
				}
			}
		})
		
	})

	$("#clear").click(function(){
		$("#taste .menuitems").css("background-color","white")
		$("#taste .menuitems").removeClass("white")
		$("#area .menuitems").css("background-color","white")
		$("#area .menuitems").removeClass("white")
		$("#condition").empty()
		var area = ""
		if ($("#condition .area")){
			area = $("#condition .area").text()
		}
		var taste = ""
		if ($("condition .taste")){
			taste = $("#condition .taste").text()
		}
		$.ajax({
			url: '/dish/getconditiondishlist',
			type: 'get',
			dataType: 'json',
			data: {"tasty":taste,"area":area},
			success:function(result){
				console.log(result)
				if (result.Ret == 200){
					$(".dishcontent").empty()
					if (result.Object != null){
					for(var i = 0;i<result.Object.length;i++){
						var text = "<a href='/dish/getdishdetail?uid="+result.Object[i].Uid+"&dishId="+result.Object[i].Id+"' class='col-sm-4 sgo' id='food'>"+
										"<div class='col-sm-12' style='padding: 0px;'><img class='img-responsive' src='"+result.Object[i].DishImg+"'></div>"+
											"<div class='col-sm-12' >"+
										"<div class='col-sm-4'>"+
										"<img src='"+result.Object[i].UserImg+"'  height='50px;' width='50px;' style='border-radius: 50px;'>"+
										"</div>"+
										"<div class='col-sm-8'>"+
										"<i class='col-sm-12'>"+result.Object[i].Name+"</i>"+
										"<i class='col-sm-12'>"+result.Object[i].DishName+"</i>"+
										"</div>"+
										"</div>"+
									"</a>"
						$(".dishcontent").append(text)
					}
				}else {
						$(".dishcontent").append(`<div class="alert alert-danger col-sm-12" role="alert">
													<strong>o(╯□╰)o</strong> 没有该搜索条件的菜谱
												</div>`)
						}
					if (result.Pref){
						$("#prev").removeClass('disabled')
					}else{	
						console.log(result.Pref)
						$("#prev").addClass('disabled')
					}
					if(result.Next){
						$("#next").removeClass('disabled')
					}else{
						$("#next").addClass('disabled')
					}
					$("#page").text(result.Page)
				}
			}
		})
	})


	$(function(){
		$("#sb").addClass('go')
	})

	$("#add").click(function(){
		console.log(123)
		$.ajax({
			url: '/user/getusermenu',
			type: 'get',
			dataType: 'json',
			success:function(result,event){
				console.log(result)
                 if (result.Ret == 200) {
                 	for (var i = 0;i< result.Object.length;i++){
                 		$("#mymenu").append("<option value="+result.Object[i].Id+">"+result.Object[i].MenuName+"</option>")
                 	}
                 	return
                 }
                 alert(result.Msg)
                 // $("#myModal").addClass('hidden')
                 location.reload()
                 return
			}
		})
	});

	$("#img").change(function(){
		imgPreview()
	})

	 function imgPreview(){
        //判断是否支持FileReader
        if (window.FileReader) {
            var reader = new FileReader();
        } else {
            alert("您的设备不支持图片预览功能，如需该功能请升级您的设备！");
        }

        //获取文件
        var file = $("#img")[0].files[0];
        var imageType = /^image\//;
        //是否是图片
        if (!imageType.test(file.type)) {
            alert("请选择图片！");
            return;
        }
        reader.readAsDataURL(file)
        //读取完成
        reader.onload = function(e) {
            //获取图片dom
            var img = document.getElementById("showimg");
            //图片路径设置为读取的图片
            img.src = e.target.result;
            console.log(img)
        };
    }



$(document).on('click', '.todelete', function(event) {
	$(this).parent().remove()
			var area = ""
		if ($("#condition .area")){
			area = $("#condition .area").text()
		}
		var taste = ""
		if ($("condition .taste")){
			taste = $("#condition .taste").text()
		}
		$.ajax({
			url: '/dish/getconditiondishlist',
			type: 'get',
			dataType: 'json',
			data: {"tasty":taste,"area":area},
			success:function(result){
				console.log(result)
				if (result.Ret == 200){
					$(".dishcontent").empty()
					if (result.Object != null){
					for(var i = 0;i<result.Object.length;i++){
						var text = "<a href='/dish/getdishdetail?uid="+result.Object[i].Uid+"&dishId="+result.Object[i].Id+"' class='col-sm-4 sgo' id='food'>"+
										"<div class='col-sm-12' style='padding: 0px;'><img class='img-responsive' src='"+result.Object[i].DishImg+"'></div>"+
											"<div class='col-sm-12' >"+
										"<div class='col-sm-4'>"+
										"<img src='"+result.Object[i].UserImg+"'  height='50px;' width='50px;' style='border-radius: 50px;'>"+
										"</div>"+
										"<div class='col-sm-8'>"+
										"<i class='col-sm-12'>"+result.Object[i].Name+"</i>"+
										"<i class='col-sm-12'>"+result.Object[i].DishName+"</i>"+
										"</div>"+
										"</div>"+
									"</a>"
						$(".dishcontent").append(text)
					}
				}else {
						$(".dishcontent").append(`<div class="alert alert-danger col-sm-12" role="alert">
													<strong>o(╯□╰)o</strong> 没有该搜索条件的菜谱
												</div>`)
						}
					if (result.Pref){
						$("#prev").removeClass('disabled')
					}else{	
						console.log(result.Pref)
						$("#prev").addClass('disabled')
					}
					if(result.Next){
						$("#next").removeClass('disabled')
					}else{
						$("#next").addClass('disabled')
					}
					$("#page").text(result.Page)
				}
			}
		})
});
