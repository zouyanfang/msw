
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
		var p = `<p class="menuitems col-sm-2 text-center taste" ><span class="glyphicon glyphicon-remove"></span>`
		p +=text+"</p>"
		$("#condition").find(".taste").remove()
		$("#condition").append(p)
	})

	$("#area .menuitems").click(function(){
		$("#area .menuitems").css("background-color","white")
		$("#area .menuitems").removeClass("white")
		var s = $(this).css("color")
		$(this).css("background-color",s)
		$(this).addClass("white")
		var text = $(this).text()
		var p = `<p class="menuitems col-sm-2 text-center area" ><span class="glyphicon glyphicon-remove"></span>`
		p +=text+"</p>"
		$("#condition").find(".area").remove()
		$("#condition").append(p)
	})

	$("#clear").click(function(){
		$("#taste .menuitems").css("background-color","white")
		$("#taste .menuitems").removeClass("white")
		$("#area .menuitems").css("background-color","white")
		$("#area .menuitems").removeClass("white")
		$("#condition").empty()
	})


	$(function(){
		$("#sb").addClass('go')
	})

	$("#add").click(function(){
		console.log("hah")
	});

	$("#file").change(function(){
		imgPreview()
		console.log("123123123edflasdjfliJWN")
	})

	 function imgPreview(){
        //判断是否支持FileReader
        if (window.FileReader) {
            var reader = new FileReader();
        } else {
            alert("您的设备不支持图片预览功能，如需该功能请升级您的设备！");
        }

        //获取文件
        var file = $("#file")[0].files[0];
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


