$(document).ready(function() {
    $("#btn_login").click(function () {
        var userName = $.trim($("#userName").val());
        var password = $.trim($("#userPassword").val());

        if (!userName) {
            $("#alertDiv").html("用户名不能为空！");
            return;
        } else if(userName.length < 5 || userName.length > 20) {
            $("#alertDiv").html("用户名在5-20字符之内");
            return;
        }

        if (!password) {
            $("#alertDiv").html("密码不能为空！");
            return;
        } else if(password.length < 6 || password.length > 16) {
            $("#alertDiv").html("密码6-16字符之内");
            return;
        }

        $("#alertDiv").html("");
        $("#form-signin").submit();


    });
});