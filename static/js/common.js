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

        }
    }

    window.sms = sms;
})()
