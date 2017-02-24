<!DOCTYPE html>
<html>
<head>
    <title>{{config "String" "globaltitle" ""}} - 服务器列表</title>
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
                <h2><i class="fa fa-server"></i>服务器列表
                    <span></span></h2>
                <div class="breadcrumbs">
                    <ol class="breadcrumb transparent-black">
                        <li>当前位置</li>
                        <li><a href="/web/public/index">首页</a></li>
                        <li class="active">服务器列表</li>
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
                                <h1><strong>服务器列表</strong></h1>
                                <span class="note">说明：标题可以排序</span>
                                <div class="controls">
                                    <a href="#" class="refresh"><i class="fa fa-refresh"></i></a>
                                    <a href="#" class="remove"><i class="fa fa-times"></i></a>
                                </div>
                            </div>
                            <div class="tile-body color transparent-black rounded-corners">
                                <div class="table-responsive">
                                    <table class="table table-datatable table-custom" id="serverListTable">
                                        <thead>
                                        <tr>
                                            <th class="sort-numeric">编号</th>
                                            <th class="sort-alpha">服务器名</th>
                                            <th class="sort-alpha">操作系统</th>
                                            <th>服务器状态</th>
                                            <th>CPU使用率</th>
                                            <th>内存使用率</th>
                                            <th>创建时间</th>
                                        </tr>
                                        </thead>
                                    </table>
                                </div>
                                <div class="modal fade" id="modalServer" tabindex="-1" role="dialog"
                                     aria-labelledby="modalServerLabel" aria-hidden="true">
                                    <div class="modal-dialog">
                                        <div class="modal-content">
                                            <div class="modal-header">
                                                <button type="button" class="close" data-dismiss="modal"
                                                        aria-hidden="true">关闭
                                                </button>
                                                <h3 class="modal-title" id="modalConfirmLabel"><strong>添加服务器</strong></h3>
                                            </div>
                                            <div class="modal-body">
                                                <form role="form" data-parsley-validate="" id="addServerForm">
                                                    <div class="form-group">
                                                        <label for="serverName">服务器名称</label>
                                                        <input type="text" class="form-control" id="serverName" required="">
                                                    </div>
                                                    <div class="form-group">
                                                        <label for="serverRoot">服务器登陆名</label>
                                                        <input type="text" class="form-control" id="serverRoot">
                                                    </div>
                                                    <div class="form-group">
                                                        <label for="serverPassword">服务器登陆密码</label>
                                                        <input type="password" class="form-control" id="serverPassword">
                                                    </div>
                                                    <div class="form-group">
                                                        <label for="IPAddress">IP地址</label>
                                                        <input type="text" class="form-control" id="IPAddress"
                                                               placeholder="IP地址">
                                                    </div>
                                                    <div class="form-group">
                                                        <label for="port">端口号</label>
                                                        <input type="text" class="form-control" id="port"
                                                               placeholder="端口号">
                                                    </div>
                                                </form>
                                            </div>
                                            <div class="modal-footer">
                                                <button class="btn btn-red" data-dismiss="modal" id="serverClose" aria-hidden="true">
                                                    关闭
                                                </button>
                                                <button class="btn btn-green" type="submit" id="addServer">添加</button>
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
<script src="/static/js/serverManager/serverList.js"></script>
</body>
</html>

