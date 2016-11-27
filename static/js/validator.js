$(function() {
    $("#userForm").bootstrapValidator({
        message: '该值无效',
        feedbackIcons: { /*input状态样式图片*/
            valid: 'glyphicon glyphicon-ok',
            invalid: 'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        live: 'enabled',
        submitButtons: 'button[type="submit"]',
        trigger: null,
        fields: {
            username: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "用户名不能为空"
                    },
                    stringLength: {
                        min: 3,
                        max: 20,
                        message: '用户名长度必须在3到20之间'
                    },
                    remote: {
                        url: "/user/",
                        message: "用户已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            return {
                                _xsrf: xsrf,
                                action:"validator",

                            }
                        },
                    },
                    regexp: {
                        regexp: /^[a-zA-Z0-9_\.]+$/,
                        message: '用户名由数字字母下划线和.组成'
                    }
                },
            },
            namezh: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "用户名(中文)不能为空"
                    },
                },
            },
            mobile: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "手机号码不能为空"
                    },
                    remote: {
                        url: "/user/",
                        message: "手机号码已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            return {
                                _xsrf: xsrf,
                                action:"validator",
                                username: $('input[name="mobile"]').val()
                            }
                        },
                    },
                    // regexp: {
                    //     regexp: /^1\d{10}$/,
                    //     message: '手机号码无效'
                    // }

                }
            },
            email: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "邮箱不能为空"
                    },
                    remote: {
                        url: "/user/",
                        message: "邮箱已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            return {
                                _xsrf: xsrf,
                                action:"validator",
                                username: $('input[name="email"]').val()
                            }
                        },
                    },
                    regexp: {
                        regexp: /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/,
                        message: '邮箱地址无效'
                    }
                }
            },
            postion: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "职位不能为空"
                    },
                }
            },
            department: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "部门不能为空"
                    },
                }
            },
            group: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "权限组不能为空"
                    },
                }
            },
            password: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "密码不能为空"
                    },
                }
            },
        },
    });
});