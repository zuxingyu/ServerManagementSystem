/**
 * Created by zuxingyu on 2016/12/10.
 */
$(document).ready(function() {

    // 将菜单栏初始化
    $('.navigation ul li').removeClass("active");
    $('.navigation ul li').removeClass("open");
    $('.navigation').addClass("collapsed");

    // 设置菜单栏选中项
    //TODO 目前只有一级目录，后面将增加二级目录判断
    var url = location.pathname;
    var urlstr = url.split("/");
    if (urlstr) {
        var lefthref = '/' + urlstr[1] + '/' + urlstr[2] + '/' + urlstr[3];
        $('.navigation a').filter(function () {
            return $(this).attr("href") == lefthref;
        }).parent().addClass("active").parent().parent().removeClass("collapsed");
    }



    // confirmation
    $('#a_logout').on('click', function () {
        $.confirm({
            title: '退出',
            content: '确定退出吗？',
            icon: 'fa fa-question-circle',
            animation: 'scale',
            closeAnimation: 'scale',
            opacity: 0.5,
            buttons: {
                'confirm': {
                    text: '确定',
                    btnClass: 'btn-info',
                    action: function () {
                        window.location.href = "/logout";
                    }
                },
                'cancel': {
                    text : '取消'
                },
            }
        });
    });
});

