$(function () {
    $(".center").fadeIn(2000)
    //登录
    $("#btn").click(function(){
        var name = $("#username")
        var password = $("#password")
        if (name.val() == ""){
            name.parent().addClass('has-warning')
            name.next().text("账户不能为空")
        }else{
            name.parent().removeClass('has-warning')
            name.next().text("")
        }
        if (password.val() == ""){
            password.parent().addClass('has-warning')
            password.next().text("密码不能为空")
        }else{
            password.parent().removeClass('has-warning')
            password.next().text("")
        }
    });

    ////失去焦点时，弹出/消除提示
    $("#username").blur(function () {
        var name = $("#username")
        if (name.val() == ""){
            name.parent().addClass('has-warning')
            name.next().text("账户不能为空")
        }else{
            name.parent().removeClass('has-warning')
            name.next().text("")
        }
    })
    $("#password").blur(function () {
        var password = $("#password")
        if (password.val() == ""){
            password.parent().addClass('has-warning')
            password.next().text("密码不能为空")
        }else{
            password.parent().removeClass('has-warning')
            $("#errors").html("")
            password.next().text("")
        }
    })

    $("#toregister").click(function(){
        $(".center").css("display","none")
        $(".centers").fadeIn(1000)
    });
    $("#login").click(function(){
        $(".centers").css("display","none")
        $(".center").fadeIn(1000)
    });

    //注册
    $("#register").click(function(){
        var name = $("#name")
        console.log(name.val())
        if (name.val() == ""){
            name.parent().addClass('has-warning')
            name.next().text("账户不能为空")
            return
        }else{
            name.parent().removeClass('has-warning')
            name.next().text("")
        }
        var password = $("#pwd")

        if (password.val() == ""){
            password.parent().addClass('has-warning')
            password.next().text("密码不能为空")
            return
        }else{
            password.parent().removeClass('has-warning')
            password.next().text("")
            return
        }
        var passwords = $("#passwords")
        if (passwords.val() != password.val()){
            passwords.parent().addClass('has-warning')
            password.next().text("密码不一致")
            return
        }else{
            passwords.parent().removeClass('has-warning')
            password.next().text("")
            return
        }

    });

    //失去焦点时，弹出/消除提示
    $("#name").blur(function () {
        var name = $("#name")
        console.log(name.val())
        if (name.val() == ""){
            name.parent().addClass('has-warning')
            name.next().text("账户不能为空")
            return
        }else{
            name.parent().removeClass('has-warning')
            name.next().text("")
        }
    });
    $("#pwd").blur(function () {
        var password = $("#pwd")

        if (password.val() == ""){
            password.parent().addClass('has-warning')
            password.next().text("密码不能为空")
            return
        }else{
            password.parent().removeClass('has-warning')
            password.next().text("")
        }
    });
    $("#passwords").blur(function () {
        var passwords = $("#passwords")
        var password = $("#pwd")
        if (passwords.val() != password.val()){
            passwords.parent().addClass('has-warning')
            password.next().text("密码不一致")
            return
        }else{
            passwords.parent().removeClass('has-warning')
            password.next().text("")
        }
    })
})
