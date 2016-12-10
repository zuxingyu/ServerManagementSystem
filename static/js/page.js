/**
 * Created by zuxingyu on 2016/12/10.
 */
$(document).ready(function() {
    $("#btn_404_refresh").click(function () {
        window.location.reload();
    });

    $("#btn_404_home").click(function () {
        window.location.href = "/web/public/index";
    });
});