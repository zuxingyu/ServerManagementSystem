/**
 * Created by Zu Xingyu on 16/11/25.
 */
(function () {

    var reSpaceCheck = /^(\d+)\.(\d+)\.(\d+)\.(\d+)(:(\d+))?$/;
    var local_or_dev = location.protocol.indexOf("file") == 0 || location.host.indexOf("localhost") == 0 || reSpaceCheck.test(location.host);

    var sms = {
        key_c_uid: "com.zuxy.sms.uid",
        key_c_nickname: "com.zuxy.sms.nickname",
        key_c_avatar: "com.zuxy.sms.avatar",
        key_c_auth: "com.zuxy.sms.authorization",
        key_c_user_type: "com.zuxy.sms.usertype",
        webServerBaseUri: local_or_dev ? "http://192.168.100.58:8082/v1" : (location.protocol + "//" + location.host + "/api"),
        // fileServerBaseUr: local_or_dev ? "http://192.168.100.58:8082/v1/file" : (location.protocol + "//" + location.host + '/v1/file'),
        count1: -1,
        // Ajax
        requestQueryServer: function (method, uri, data, success, fail) {
            var requestUri = kp.webServerBaseUri + uri;
            $.ajax({
                type: method,
                url: requestUri,
                data: data,
                contentType: "application/x-www-form-urlencoded",
                success: function (data) {
                    if (success) {
                        success && success(data);
                    }
                },
                beforeSend: function (xhr) {
                    if (authorization) {
                        xhr.setRequestHeader("Authorization", authorization);
                    }
                },
                error: function (data) {
                    console.log(data);
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })
        },
        requestFormServer: function (method, uri, data, success, fail) {
            var requestUri = kp.webServerBaseUri + uri;
            $.ajax({
                type: method,
                url: requestUri,
                data: data,
                contentType: "application/x-www-form-urlencoded",
                success: function (data) {
                    if (success) {
                        success && success(data);
                    }
                },
                beforeSend: function (xhr) {
                    if (authorization) {
                        xhr.setRequestHeader("Authorization", authorization);
                    }
                },
                error: function (data) {
                    console.log(data);
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })
        },
        // Ajax JSON
        requestJsonServer: function (method, uri, data, success, fail) {
            var requestUri = kp.webServerBaseUri + uri;
            $.ajax({
                type: method,
                url: requestUri,
                data: JSON.stringify(data),
                dataType: "json",
                contentType: "application/json;charset=UTF-8",
                success: function (data) {
                    if (success) {
                        success && success(data);
                    }
                },
                beforeSend: function (xhr) {
                    if (authorization) {
                        xhr.setRequestHeader("Authorization", authorization);
                    }
                },
                error: function (data) {
                    console.log(data);
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })

        },
        windowSizeController: function () {
            var i = $(window).width(), _width;
            if (i <= 320) {
                _width = 320;
            } else if (i <= 480) {
                _width = i;
            } else {
                _width = i - $('#sidebar').width();
            }
            var h = $(window).height();
            $('#wrapper').css('height', h - 70);
            $('.content').css('height', h - 70);
            $('#main').css('width', _width);
        },
        // show Page
        showPage: function (rescount, pageindex, pagenum, pagetime, callback) {
            this.showPageTag('pagenum', rescount, pageindex, pagenum, pagetime, callback);
        },
        showPage2: function (rescount, pageindex, pagenum, pagetime, callback) {
            this.showPageTag2('pagenum2', rescount, pageindex, pagenum, pagetime, callback);
        },
        showPageTag: function (pagenumtag, rescount, pageindex, pagenum, pagetime, callback) {
            var that = this;
            var htmls = [];

            var currentPage = pageindex;
            var tempage = parseInt(rescount / pagenum);
            if (rescount % pagenum > 0) {
                tempage += 1;
            }
            pagecount = tempage;
            var pages = pagecount;

            if (currentPage <= 8) {

                for (i = 1; i <= (pages <= 10 ? pages : 10); i++) {

                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            } else if (currentPage > 8 && pages <= (~~currentPage + 5)) {
                for (i = pages - 8; i <= pages; i++) {
                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            } else if (currentPage > 8 && pages > (~~currentPage + 5)) {

                for (i = (~~currentPage - 4); i <= (~~currentPage + 5); i++) {
                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            }
            that.bindpageTag(pagenumtag, htmls, pagetime, callback);
        },
        showPageTag2: function (pagenumtag, rescount, pageindex, pagenum, pagetime, callback) {
            var that = this;
            var htmls = [];

            var currentPage = pageindex;
            var tempage = parseInt(rescount / pagenum);
            if (rescount % pagenum > 0) {
                tempage += 1;
            }
            pagecount = tempage;
            var pages = pagecount;

            if (currentPage <= 8) {

                for (i = 1; i <= (pages <= 10 ? pages : 10); i++) {

                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            } else if (currentPage > 8 && pages <= (~~currentPage + 5)) {
                for (i = pages - 8; i <= pages; i++) {
                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            } else if (currentPage > 8 && pages > (~~currentPage + 5)) {

                for (i = (~~currentPage - 4); i <= (~~currentPage + 5); i++) {
                    if (i == currentPage) {
                        var temple = "<a href='javascript:void(0);' class='curr' data-num='" + i
                            + "'>" + i + "</a>";
                    }
                    else {
                        var temple = "<a href='javascript:void(0);' data-num='" + i + "'>" + i
                            + "</a>";
                    }
                    htmls.push(temple);
                }

            }
            that.bindpageTag2(pagenumtag, htmls, pagetime, callback);
        },
        // Bind Page
        bindpageTag: function (pagenumtag, htmls, pagetime, callback) {
            $("#" + pagenumtag + "").html(htmls.join(''));
            $("#" + pagenumtag + " a").each(function () {
                $(this).click(function () {
                    var num = $(this).attr("data-num");
                    $("#pagenum a").removeClass("curr");
                    $(this).addClass("curr");
                    callback(num, pagetime);
                });
            })
        },
        bindpageTag2: function (pagenumtag, htmls, pagetime, callback) {
            $("#" + pagenumtag + "").html(htmls.join(''));
            $("#" + pagenumtag + " a").each(function () {
                $(this).click(function () {
                    var num = $(this).attr("data-num");
                    $("#pagenum2 a").removeClass("curr");
                    $(this).addClass("curr");
                    callback(num, pagetime);
                });
            })
        },
        bindpage: function (htmls, pagetime, callback) {
            this.bindpageTag("pagenum", htmls, pagetime, callback);
            //$("#pagenum").html(htmls.join(''));
            //$("#pagenum a").each(function () {
            //    $(this).click(function () {
            //        var num = $(this).attr("data-num");
            //        $("#pagenum a").removeClass("curr");
            //        $(this).addClass("curr");
            //        callback(num, pagetime);
            //    });
            //})
        },
        getUrlParam: function (name, url) {
            var str = url || window.location.search || window.location.hash,
                result = null;
            if (!name || str === '') {
                return result;
            }
            result = str.match(new RegExp('(^|&|[\?#])' + name + '=([^&]*)(&|$)'));
            return result === null ? result : decodeURIComponent(result[2]);
        },
        getFileUrl: function (uri) {
            if (Boolean(uri)) {
                if (uri.indexOf("http") == -1) {
                    if (uri.substring(0, 1) === "/") {
                        return kp.fileServerBaseUr + uri;
                    } else {
                        return kp.fileServerBaseUr + "/" + uri;
                    }
                } else {
                    return uri;
                }
            }
            else {
                return "";
            }
        },
        getAvatar: function (uri, gender) {
            if (uri) {
                return kp.getFileUrl(uri);
            } else {
                if (gender == 1) {
                    return "images/default_boy.png"
                } else {
                    return "images/default_girl.png"
                }
            }
        },
        addMenu: function () {

            var menustr = "";
            menustr += '<div>';
            menustr += '<h4 ><span _data="one">用户管理</span></h4>';
            menustr += '<dl class="one"><dt><a href="indentify.html" class="menu" >认证审核</a></dt>';
            menustr += '<dt><a href="userinfo.html" class="menu">用户信息</a></dt>';
            menustr += '<dt><a href="usermanager.html" class="menu">账户管理</a></dt>';
            menustr += '<dt><a href="account.html" class="menu">虚拟账号</a></dt></dl>';

            menustr += '<h4 ><span _data="two">内容管理</span></h4>';
            menustr += '<dl class="two"><dt><a href="bannerlist.html" class="menu">Banner</a></dt>';
            menustr += '<dt><a href="gift.html" class="menu">礼物管理</a></dt>';
            menustr += '<dt><a href="videomanager.html" class="menu">直播监控</a></dt>';
            menustr += '<dt><a href="report.html" class="menu">举报管理</a></dt>';
            menustr += '<dt><a href="activity.html" class="menu">抽奖活动管理</a></dt>';

            menustr += '<dt><a href="livevideo.html" class="menu">直播回放</a></dt>';
            menustr += '<dt><a href="audiovisual.html" class="menu">视听</a></dt>';

            menustr += '<dt><a href="notice.html" class="menu">意见反馈/官方通知</a></dt></dl>';

            menustr += '<dt><a href="justicelist.html" class="menu">正义帮列表</a></dt></dl>';
            menustr += '<dt><a href="helplist.html" class="menu">帮助中心</a></dt></dl>';

            menustr += '</div>';

            $("#menubox").html(menustr);
            var curl = window.location.href;
            $("#menubox a").each(function () {
                if (curl.indexOf($(this).attr("href")) > -1) {
                    $(this).addClass("current");
                } else {
                    $(this).removeClass("current");
                }
            })

            $("#menubox span").click(function (e) {
                //alert(e.target)
                if (e.target == this) {
                    var target = $(this).attr("_data");
                    $("." + target).toggleClass('hidden');
                }
            })


        },
        exit: function () {
            $.cookie(kp.key_c_uid, null);
            $.cookie(kp.key_c_auth, null);
            var curl = window.location.href;
            if (curl.indexOf("login.html") <= -1) {
                location.href = "login.html";
            }
        },
        checkNormalAdmin: function () {
            var userType = $.cookie(kp.key_c_user_type);
            return (userType & 1 << 2) == 1 << 2;
        },
        checkSuperAdmin: function () {
            var userType = $.cookie(kp.key_c_user_type);
            return (userType & 1 << 1) == 1 << 1;
        },
        getdata: function (method, uri, data, success, fail) {
            var requestUri = uri
            $.getJSON(requestUri, function (result) {
                if (success) {
                    success && success(result);
                }
            });
        },
        GetQueryStringlocation: function (name) {
            var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
            var r = window.location.search.substr(1).match(reg);
            if (r != null)return unescape(r[2]);
            return null;
        },
        isPositiveInt: function (val) {
            if ((/^(\+|-)?\d+$/.test(val)) && val > 0) {
                return true;
            } else {
                return false;
            }
        }
    }

    var authorization = $.cookie(kp.key_c_auth);
    window.kp = kp;
    if (Boolean(window.login) == false) {
        if (Boolean(authorization) == false) {
            location.href = "login.html"
        }
    }
})()

//window.onload = function() {
//    var data = {type: 2};
//    kp.requestQueryServer("GET", "/real-name-auth/total", data, function (res) {
//        if (res.header.code == "1000") {
//            $(".checkcount").html(res.data.total);
//        }
//    })
//};
