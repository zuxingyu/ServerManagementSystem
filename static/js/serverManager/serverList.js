/**
 * Created by zuxingyu on 2017/1/12.
 */
$(function () {
    $.fn.dataTableExt.oStdClasses.sPaging = 'dataTables_paginate paging_bootstrap paging_custom';
    /* Define two custom functions (asc and desc) for string sorting */
    jQuery.fn.dataTableExt.oSort['string-case-asc']  = function(x,y) {
        return ((x < y) ? -1 : ((x > y) ?  1 : 0));
    };

    jQuery.fn.dataTableExt.oSort['string-case-desc'] = function(x,y) {
        return ((x < y) ?  1 : ((x > y) ? -1 : 0));
    };

    var addRowLink = '<a href="#modalServer" id="addRow" class="btn btn-green add-row" data-toggle="modal">添加服务器</a>';
    $('#serverListTable').append(addRowLink);

    $('#serverListTable').dataTable({
        "aaSorting": [
            [0, "desc"]
        ],
        "aLengthMenu": [[10, 25, 50], ["10条", "25条", "50条"]],//二组数组，第一组数量，第二组说明文字;
        "bPaginate": true, //翻页功能
        "bLengthChange": true, //改变每页显示数据数量
        "bSort": true, //排序功能
        "bFilter": true, //过滤功能
        "bInfo": true,//页脚信息
        "bAutoWidth": true,//自动宽度
        "bStateSave": false,//保存上次操作
        "bProcessing": true,//显示正在处理
        "bServerSide": true,//延迟加载
        "sDom": "R<'row'<'col-md-6'l><'col-md-6'f>r>" +
        "t" +
        "<'row'<'col-md-4 sm-center'i><'col-md-4'><'col-md-4 text-right sm-center'p>>",
        "oLanguage": {  //本地化设置
            "sLengthMenu": "每页显示 _MENU_ 条记录",
            "sZeroRecords": "抱歉， 没有找到",
            "sInfo": "从 _START_ 到 _END_ /共 _TOTAL_ 条数据",
            "sInfoEmpty": "没有数据",
            "sInfoFiltered": "(从 _MAX_ 条数据中检索)",
            // "sProcessing": "正在加载数据...",
            "sProcessing": "<img src='/static/assets/images/loader.gif' />", //这里是给服务器发请求后到等待时间显示的 加载gif
            "oPaginate": {
                "sFirst": "首页",
                "sPrevious": "前一页",
                "sNext": "后一页",
                "sLast": "尾页",
            },
            "sZeroRecords": "没有搜索到数据",
            "sSearch": "",
        },
        "sAjaxSource": "/web/serverManager/listData",
        "aoColumns": [ //这个属性下的设置会应用到所有列，按顺序没有是空
            {"mData": 'id'}, //mData 表示发请求时候本列的列明，返回的数据中相同下标名字的数据会填充到这一列
            {"mData": 'aliasName'},
            {"mData": 'serverOS'},
            {"mData": 'serverState'},
            {"mData": 'cpuRate'},
            {"mData": 'memoryRate'},
            {"mData": 'Created'},
        ],
        "aoColumnDefs": [{"bSortable": false, "aTargets": [4, 5]}, {"bSearchable": false, "aTargets": [ 2, 3, 4, 5]}],
        "iDisplayLength": 10,//用于指定一屏显示的条数，需开启分页器
        "fnRowCallback": function (nRow, aData, iDisplayIndex) {// 当创建了行，但还未绘制到屏幕上的时候调用，通常用于改变行的class风格
            if (aData.serverState == 0) {
                $('td:eq(3)', nRow).html("<span class='label label-success'>正常</span>");
            } else if (aData.serverState == 1) {
                $('td:eq(3)', nRow).html("<span class='label label-redbrown'>停用</span>");
            }

            var classStr = "";
            if(aData.cpuRate >= 70 && aData.cpuRate < 90){
                classStr = "progress-bar-orange";
            } else if(aData.cpuRate > 90){
                classStr = "progress-bar-red";
            } else {
                classStr = "progress-bar-green";
            }

            var cpuRateStr = "<div class='progress progress-striped active'>"+
                "<div class='progress-bar "+classStr+"' role='progressbar' aria-valuenow='" + aData.cpuRate + "' aria-valuemin='0' aria-valuemax='100' style='width: " + aData.cpuRate + "%'>" +
                aData.cpuRate + "%" +
                "</div>"+
                "</div>";

            $('td:eq(4)', nRow).html(cpuRateStr);

            if(aData.memoryRate >= 70 && aData.memoryRate < 90){
                classStr = "progress-bar-orange";
            } else if(aData.memoryRate > 90){
                classStr = "progress-bar-red";
            } else {
                classStr = "progress-bar-green";
            }

            var memoryRateStr = "<div class='progress progress-striped active'>"+
                "<div class='progress-bar "+classStr+"' role='progressbar' aria-valuenow='" + aData.memoryRate + "' aria-valuemin='0' aria-valuemax='100' style='width: " + aData.memoryRate + "%'>" +
                aData.memoryRate + "%" +
                "</div>"+
                "</div>";
            $('td:eq(5)', nRow).html(memoryRateStr);

            $('td:eq(6)', nRow).html(sms.UnixToDate(aData.Created * 1000));
            return nRow;
        },
        "asStripClasses": ["odd", "even"],
        "fnInitComplete": function (oSettings, json) {
            $('.dataTables_filter input').attr("placeholder", "搜索服务器名");
        }
    });
    $('.dataTables_length select').chosen({disable_search_threshold: 10});

    $('#serverClose').click(function () {
        $("#serverName").val("");
        $("#serverRoot").val("");
        $("#serverPassword").val("");
        $("#IPAddress").val("");
        $("#port").val("");
    });
    $('#addServer').click(function () {
        alert("----");
    });

    // $("#addServerForm").parsley();
})