<!DOCTYPE html>
<html>
<head>
    <title>{{config "String" "globaltitle" ""}} - 用户列表</title>
    {{template "template/meta.tpl" .}}
    <link rel="stylesheet" href="/static/assets/js/vendor/chosen/css/chosen.min.css">
    <link rel="stylesheet" href="/static/assets/js/vendor/chosen/css/chosen-bootstrap.css">
    <link rel="stylesheet" href="/static/assets/js/vendor/datatables/css/dataTables.bootstrap.css">
    <link rel="stylesheet" href="/static/assets/js/vendor/datatables/css/ColVis.css">
    <link rel="stylesheet" href="/static/assets/js/vendor/datatables/css/TableTools.css">
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
                <h2><i class="fa fa-user"></i>用户列表
                    <span></span></h2>
                <div class="breadcrumbs">
                    <ol class="breadcrumb transparent-black">
                        <li>当前位置</li>
                        <li><a href="/web/public/index">首页</a></li>
                        <li class="active">用户列表</li>
                    </ol>
                </div>
            </div>
            <!-- /page header -->
            <!-- content main container -->
            <div class="main">
                <!-- row -->
                <div class="row">
                    <!-- col 6 -->
                    <div class="col-md-12">
                        <section class="tile transparent">
                            <div class="tile-header transparent">
                                <h1><strong>用户列表</strong></h1>
                                <span class="note">说明：标题可以排序</span>
                                <div class="controls">
                                    <a href="#" class="refresh"><i class="fa fa-refresh"></i></a>
                                    <a href="#" class="remove"><i class="fa fa-times"></i></a>
                                </div>
                            </div>
                            <div class="tile-body color transparent-black rounded-corners">
                                <div class="table-responsive">
                                    <table class="table table-datatable table-custom" id="userListTable">
                                        <thead>
                                            <tr>
                                                <th class="sort-numeric">编号</th>
                                                <th class="sort-alpha">昵称</th>
                                                <th class="sort-amount">用户名</th>
                                                <th class="sort-numeric">手机号</th>
                                                <th>注册时间</th>
                                                <th>用户状态</th>
                                                <th>操作</th>
                                            </tr>
                                        </thead>
                                    </table>
                                </div>
                                <div class="modal fade" id="modalConfirm" tabindex="-1" role="dialog" aria-labelledby="modalConfirmLabel" aria-hidden="true">
                                    <div class="modal-dialog">
                                        <div class="modal-content">
                                            <div class="modal-header">
                                                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">Close</button>
                                                <h3 class="modal-title" id="modalConfirmLabel"><strong>Modal</strong> title</h3>
                                            </div>
                                            <div class="modal-body">
                                                <form role="form">

                                                    <div class="form-group">
                                                        <label for="exampleInput">Normal input field</label>
                                                        <input type="text" class="form-control" id="exampleInput">
                                                    </div>

                                                    <div class="form-group">
                                                        <label for="passwordInput">Password input field</label>
                                                        <input type="password" class="form-control" id="passwordInput">
                                                    </div>

                                                    <div class="form-group">
                                                        <label for="placeholderInput">Input with placeholder</label>
                                                        <input type="text" class="form-control" id="placeholderInput" placeholder="This is a placeholder...">
                                                    </div>

                                                    <div class="form-group">
                                                        <label>Normal textarea</label>
                                                        <textarea class="form-control" rows="3"></textarea>
                                                    </div>

                                                </form>
                                            </div>
                                            <div class="modal-footer">
                                                <button class="btn btn-red" data-dismiss="modal" aria-hidden="true">Close</button>
                                                <button class="btn btn-green">Save changes</button>
                                            </div>
                                        </div><!-- /.modal-content -->
                                    </div><!-- /.modal-dialog -->
                                </div><!-- /.modal -->
                            </div>
                        </section>
                    </div>
                </div>
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
<script src="/static/assets/js/vendor/datatables/jquery.dataTables.min.js"></script>
<script src="/static/assets/js/vendor/datatables/ColReorderWithResize.js"></script>
<script src="/static/assets/js/vendor/datatables/colvis/dataTables.colVis.min.js"></script>
<script src="/static/assets/js/vendor/datatables/tabletools/ZeroClipboard.js"></script>
<script src="/static/assets/js/vendor/datatables/tabletools/dataTables.tableTools.min.js"></script>
<script src="/static/assets/js/vendor/datatables/dataTables.bootstrap.js"></script>
<script src="/static/js/userManager/userList.js"></script>
</body>
</html>

