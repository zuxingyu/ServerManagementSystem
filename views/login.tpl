<!DOCTYPE html>
<html>
<head>
    <title>服务器监控系统-登录界面</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8"/>
    <link rel="icon" type="image/ico" href="/static/images/icon/favicon.ico" />
    <link href="/static/assets/css/vendor/bootstrap/bootstrap.min.css" rel="stylesheet" />
    <link href="/static/assets/css/font-awesome.min.css" rel="stylesheet" />
    <link href="/static/assets/css/vendor/bootstrap-checkbox.css" rel="stylesheet" />
    <link href="/static/assets/css/minimal.css" rel="stylesheet" />
    <link href="/static/css/login.css" rel="stylesheet"/>
</head>
<body class="bg-1">


<!-- Wrap all page content here -->
<div id="wrap">
    <!-- Make page fluid -->
    <div class="row">
        <!-- Page content -->
        <div id="content" class="col-md-12 full-page login">
            <div class="inside-block">
                <img src="static/assets/images/logo-big.png" alt class="logo">
                <h1><strong>服务器监控系统</strong></h1>
                <h5>By Zu Xingyu</h5>
                <div id="alertDiv">{{.error}}</div>
                <form id="form-signin" class="form-signin form-horizontal" action="/login" method="post">
                    <section>
                        <div class="input-group">
                            <input type="text" id="userName" class="form-control" name="username" required="" placeholder="用户名"/>
                            <div class="input-group-addon"><i class="fa fa-user"></i></div>
                        </div>
                        <div class="input-group">
                            <input type="password" id="userPassword" class="form-control" required="" name="password" placeholder="密码"/>
                            <div class="input-group-addon"><i class="fa fa-key"></i></div>
                        </div>
                    </section>
                    <section class="controls">
                        <div class="checkbox check-transparent">
                            <input type="checkbox" value="1" id="remember" checked>
                            <label for="remember">记住我</label>
                            <a href="#">忘记密码?</a>
                        </div>
                    </section>
                    <section class="log-in">
                        <button type="button" id="btn_login" class="btn btn-greensea">登录</button>
                        <!--<span>or</span>-->
                        <!--<button class="btn btn-slategray">Create an account</button>-->
                    </section>
                </form>
            </div>
        </div>
        <!-- /Page content -->
    </div>
</div>
<!-- Wrap all page content end -->
<script type="text/javascript" src="/static/assets/js/jquery.js"></script>
<script type="text/javascript" src="/static/assets/js/vendor/bootstrap/bootstrap.min.js"></script>
<script type="text/javascript" src="/static/js/login.js"></script>
</body>
</html>

