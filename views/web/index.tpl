<!DOCTYPE html>
<html>
<head>
    <title>{{config "String" "globaltitle" ""}}</title>
    {{template "template/meta.tpl" .}}
</head>
<body class="bg-1">
<!-- Preloader 加载 -->
<div class="mask">
    <div id="loader"></div>
</div>
<!--/Preloader -->

<!-- Wrap all page content here -->
<div id="wrap">
    <!-- Make page fluid -->
    <div class="row">
        {{template "template/leftAndhead.tpl" .}}
        <!-- Page content -->
        <div id="content" class="col-md-12">
            <!-- page header -->
            <div class="pageheader">
                <h2><i class="fa fa-home"></i> 首页
                    <span></span></h2>
                <div class="breadcrumbs">
                    <ol class="breadcrumb transparent-black">
                        <li>当前位置</li>
                        <li>首页</li>
                    </ol>
                </div>
            </div>
            <!-- /page header -->
            <!-- content main container -->
            <div class="main">
                这是一个首页
            </div>
            <!-- /content container -->
        </div>
        <!-- Page content end -->
        {{template "template/message.tpl" .}}
    </div>
    <!-- Make page fluid-->
</div>
<!-- Wrap all page content end -->
{{template "template/script.tpl" .}}
</body>
</html>

