/**
 * Created by zuxingyu on 2016/12/17.
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

    $('#userListTable').dataTable({
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
        "sAjaxSource": "/web/userManager/listData",
        "aoColumns": [ //这个属性下的设置会应用到所有列，按顺序没有是空
            {"mData": 'id'}, //mData 表示发请求时候本列的列明，返回的数据中相同下标名字的数据会填充到这一列
            {"mData": 'nickName'},
            {"mData": 'userName'},
            {"mData": 'mobile'},
            {"mData": 'Created'},
            {"mData": 'userState'},
            {"sDefaultContent": ''}, // sDefaultContent 如果这一列不需要填充数据用这个属性，值可以不写，起占位作用
        ],
        "aoColumnDefs": [{"bSortable": false, "aTargets": [6]}, {"bSearchable": false, "aTargets": [ 4, 5, 6]}],
        "iDisplayLength": 10,//用于指定一屏显示的条数，需开启分页器
        "fnRowCallback": function (nRow, aData, iDisplayIndex) {// 当创建了行，但还未绘制到屏幕上的时候调用，通常用于改变行的class风格
            if (aData.userState == 1) {
                $('td:eq(6)', nRow).html("<a></a>");
                $('td:eq(5)', nRow).html("正常");
            } else if (aData.userState == 0) {
                $('td:eq(5)', nRow).html("停用");
            }
            $('td:eq(4)', nRow).html(sms.UnixToDate(aData.Created * 1000));
            return nRow;
        },
        "asStripClasses": ["odd", "even"],
        "fnInitComplete": function (oSettings, json) {
            $('.dataTables_filter input').attr("placeholder", "搜索");
        }
    });
    $('.dataTables_length select').chosen({disable_search_threshold: 10});

})