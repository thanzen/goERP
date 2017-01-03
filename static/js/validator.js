$(function() {
    //用户
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
            name: {
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
                                action: "validator",

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
                                action: "validator",
                                name: $('input[name="mobile"]').val()
                            }
                        },
                    },
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
                                action: "validator",
                                name: $('input[name="email"]').val()
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
    //产品分类
    $("#productCategoryForm").bootstrapValidator({
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
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "产品类别不能为空"
                    },
                    remote: {
                        url: "/product/category/",
                        message: "该类别已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            var recordId = $("input[name ='_recordId']");
                            res = {
                                _xsrf: xsrf,
                                action: "validator",
                            }
                            if (recordId != undefined && recordId[0]) {
                                recordId = recordId[0].value;
                                res.recordId = recordId;
                            }
                            return res
                        },
                    },
                },
            },
        },
    });
    //产品属性
    $("#productAttributeForm").bootstrapValidator({
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
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "属性名称不能为空"
                    },
                    remote: {
                        url: "/product/attribute/",
                        message: "该属性名称已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            var recordId = $("input[name ='_recordId']");
                            res = {
                                _xsrf: xsrf,
                                action: "validator",
                            }
                            if (recordId != undefined && recordId[0]) {
                                recordId = recordId[0].value;
                                res.recordId = recordId;
                            }
                            return res
                        },
                    },
                },
            },
        },
    });
    //产品属性值
    $("#productAttributeValueForm").bootstrapValidator({
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
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "属性值不能为空"
                    },
                    remote: {
                        url: "/product/attributevalue/",
                        message: "该属性值已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            var recordId = $("input[name ='_recordId']");
                            var attributeId = $("select[name='productAttributeID']");
                            res = {
                                _xsrf: xsrf,
                                action: "validator",
                            }
                            if (recordId != undefined && recordId[0]) {
                                recordId = recordId[0].value;
                                res.recordId = recordId;
                            }
                            if (attributeId != undefined && attributeId[0]) {
                                attributeId = attributeId[0].value;
                                res.attributeId = attributeId;
                            }
                            return res
                        },
                    },
                },
            },
        },
    });
    //计量单位分类
    $("#productUomCategForm").bootstrapValidator({
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
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "计量单位分类不能为空"
                    },
                    remote: {
                        url: "/product/uomcateg/",
                        message: "该计量单位分类已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            var recordId = $("input[name ='_recordId']");
                            res = {
                                _xsrf: xsrf,
                                action: "validator",
                            }
                            if (recordId != undefined && recordId[0]) {
                                recordId = recordId[0].value;
                                res.recordId = recordId;
                            }
                            return res
                        },
                    },
                },
            },
        },
    });
    //计量单位分类
    $("#productUomForm").bootstrapValidator({
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
            category: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "计量单位分类不能为空"
                    },
                }
            },
            name: {
                message: "该值无效",
                validators: {
                    notEmpty: {
                        message: "计量单位名称不能为空"
                    },
                    remote: {
                        url: "/product/uom/",
                        message: "该计量单位名称已经存在",
                        dataType: "json",
                        delay: 200,
                        type: "POST",
                        data: function() {
                            var xsrf = $("input[name ='_xsrf']")[0].value;
                            var recordId = $("input[name ='_recordId']");
                            res = {
                                _xsrf: xsrf,
                                action: "validator",
                            }
                            if (recordId != undefined && recordId[0]) {
                                recordId = recordId[0].value;
                                res.recordId = recordId;
                            }
                            return res
                        },
                    },
                },
            },
        },
    });
});