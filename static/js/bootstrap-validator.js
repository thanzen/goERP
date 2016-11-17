$(function() {
    $("#userForm").bootstrapValidator({
        message: '该值无效',
        feedbackIcons: { /*input状态样式图片*/
            valid: 'glyphicon glyphicon-ok',
            invalid: 'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "用户名不能为空"
                    },
                    stringLength: {
                        min: 5,
                        max: 30,
                        message: '用户名长度必须在6到30之间'
                    },
                    threshold: 5,
                    remote: {
                        url: "/user",
                        message: "用户已经存在",
                        delay: 2000,
                        type: "POST",
                        data: function() {
                            
                        },
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '用户名由数字字母下划线和.组成'
                    }
                },
            }
        },
    });
});