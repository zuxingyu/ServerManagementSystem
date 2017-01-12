/**
 * Created by Zu Xingyu on 16/11/25.
 */
(function () {

    var reSpaceCheck = /^(\d+)\.(\d+)\.(\d+)\.(\d+)(:(\d+))?$/;
    var local_or_dev = location.protocol.indexOf("file") == 0 || location.host.indexOf("localhost") == 0 || reSpaceCheck.test(location.host);

    var sms = {
        webServerBaseUri: location.protocol + "//" + location.host + ":8080",
        // Ajax
        requestQueryServer: function (method, uri, data, success, fail) {
            var requestUri = sms.webServerBaseUri + uri;
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
                error: function (data) {
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })
        },
        requestFormServer: function (method, uri, data, success, fail) {
            var requestUri = sms.webServerBaseUri + uri;
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
                error: function (data) {
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })
        },
        // Ajax JSON
        requestJsonServer: function (method, uri, data, success, fail) {
            var requestUri = sms.webServerBaseUri + uri;
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
                error: function (data) {
                    if (fail) {
                        fail && fail(data);
                    }
                }
            })
        },
        UnixToDate: function (unixTime, isFull, timeZone) {
            if (typeof (timeZone) == 'number') {
                unixTime = parseInt(unixTime) + parseInt(timeZone) * 60 * 60;
            }
            var time = new Date(unixTime);
            var ymdhis = "";
            ymdhis += time.getFullYear() + "/";
            ymdhis +=
                (time.getMonth() + 1) < 10 ? "0" + (time.getMonth() + 1) + "/"
                    : (time.getMonth() + 1) + "/";
            ymdhis += time.getDate() < 10 ? "0" + time.getDate() : time.getDate();
            if (isFull === true) {
                ymdhis +=
                    time.getHours() < 10 ? " 0" + time.getHours() + ":" : " "
                    + time.getHours()
                    + ":";
                ymdhis +=
                    time.getMinutes() < 10 ? "0" + time.getMinutes() + ":" : ""
                    + time.getMinutes()
                    + ":";
                ymdhis +=
                    time.getSeconds() < 10 ? "0" + time.getSeconds() : time.getSeconds();
            }
            return ymdhis;
        },
    }

    window.sms = sms;
})()
