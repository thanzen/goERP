$(document).ready(function() {

    $.extend($.fn.bootstrapTable.defaults, {
        method: "post",
        dataType: "json",
        locale: "zh-CN",
        contentType: "application/x-www-form-urlencoded",
        sidePagination: "server",
        stickyHeader: true, //表头固定
        stickyHeaderOffsetY: (function() {
            var stickyHeaderOffsetY = 0;
            if ($('.navbar-fixed-top').css('height')) {
                stickyHeaderOffsetY = +$('.navbar-fixed-top').css('height').replace('px', '');
            }
            if ($('.navbar-fixed-top').css('margin-bottom')) {
                stickyHeaderOffsetY += +$('.navbar-fixed-top').css('margin-bottom').replace('px', '');
            }
            return stickyHeaderOffsetY + 'px';
        })(), //设置偏移量
        dataField: "data",
        pagination: true,
        pageNumber: 1,
        pageSize: 10,
        pageList: [10, 25, 50, 100, 500, 1000],
        // onClickRow: function(row, $element) {
        //     //$element是当前tr的jquery对象
        //     $element.css("background-color", "green");
        // },//单击row事件
    });
    var displayTable = function(selectId, ajaxUrl, columns, onExpandRow) {
            var $tableNode = $(selectId);
            var options = {
                url: ajaxUrl,
                queryParams: function(params) {
                    var xsrf = $("input[name ='_xsrf']");
                    if (xsrf != undefined) {
                        params._xsrf = xsrf[0].value;
                    }
                    params.action = 'table';
                    return params;
                },
                columns: columns,
            }
            if (onExpandRow != undefined) {
                options.detailView = true;
                options.onExpandRow = onExpandRow;
            }
            $tableNode.bootstrapTable(options);
        }
        //用户表
    displayTable("#table-user", "/user/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc" },
        { title: "中文名称", field: 'namezh', sortable: true, order: "desc" },
        { title: "部门", field: 'department', sortable: true, order: "desc" },
        { title: "职位", field: 'position', sortable: true, order: "desc" },
        { title: "邮箱", field: 'email', sortable: true, order: "desc" },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc" },
        { title: "座机", field: 'tel', sortable: true, order: "desc" },
        { title: "QQ", field: 'qq', sortable: true, order: "desc" },
        { title: "微信", field: 'wechat', sortable: true, order: "desc" },
        { title: "管理员", field: 'isadmin', sortable: true, order: "desc" },
        { title: "有效", field: 'active', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                if (row.active == "有效") {
                    html += "<a href='/user/" + row.Id + "?action=invalid' class='table-action btn btn-xs btn-danger'>无效&nbsp<i class='fa fa-close'></i></a>";
                } else {
                    html += "<a href='/user/" + row.Id + "?action=active' class='table-action btn btn-xs btn-success'>有效&nbsp<i class='fa fa-check'></i></a>";
                }
                html += "<a href='/user/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/user/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }

    ], function(index, row, $detail) {
        console.log(row);
        var html = "1231231";
        var params = (function() {
            var params = {};
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action = 'table';
            params.offset = 0;
            params.limit = 5;
            return params;
        })();
        $.ajax({
            url: "/user/",
            dataType: "json",
            type: "POST",
            async: false,
            data: params,
            success: function(data) {
                console.log(data);
                html = "ok";
                $detail.html(data.total);
            },
            error: function(error) {

                html = error;
                $detail.html(html);
            },

        });


    });
    //登录记录表
    displayTable("#table-record", "/record/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc" },
        { title: "邮箱", field: 'email', sortable: true, order: "desc" },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc" },
        { title: "开始时间", field: 'start_time', sortable: true, order: "desc" },
        { title: "结束时间", field: 'end_time', sortable: true, order: "desc" },
        { title: "IP地址", field: 'ip', sortable: true, order: "desc" },
        { title: "用户代理", field: "UserAgent" }

    ]);
    //国家表
    displayTable("#table-country", "/address/country/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "国家", field: 'name', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/address/country/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/address/country/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);
    //省份表
    displayTable("#table-province", "/address/province/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "省份", field: 'name', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/address/province/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/address/province/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //城市表
    displayTable("#table-city", "/address/city/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "城市", field: 'name', sortable: true, order: "desc" },
        { title: "省份", field: 'province', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/address/city/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/address/city/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);
    //区县表
    displayTable("#table-district", "/address/district/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "地区", field: 'name', sortable: true, order: "desc" },
        { title: "城市", field: 'city', sortable: true, order: "desc" },
        { title: "省份", field: 'province', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/address/district/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/address/district/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //产品属性
    displayTable("#table-product-attribute", "/product/attribute/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "属性名", field: 'name', sortable: true, order: "desc" },
        { title: "属性编码", field: 'code', sortable: true, order: "desc" },
        { title: "属性序号", field: 'sequence', sortable: true, order: "desc" },
        {
            title: "属性值",
            field: 'childs',
            align: "center",
            formatter: function cellStyle(value, row, index) {
                var datas = row.values;
                var html = "";
                for (key in datas) {
                    html += "<span  class='display-block label label-primary'>" + datas[key] + "</span>";
                }
                return html;
            }
        },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/product/attribute/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/product/attribute/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //产品类别
    displayTable("#table-product-category", "/product/category/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "类别名", field: 'name', sortable: true, order: "desc" },
        { title: "上级", field: 'parent', sortable: true, order: "desc" },
        { title: "上级路径", field: 'path', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/product/category/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/product/category/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //产品款式
    displayTable("#table-product-template", "/product/template/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "款式编码", field: 'defaultCode', sortable: true, order: "desc" },
        { title: "款式类别", field: 'category', sortable: true, order: "desc" },
        { title: "产品款式", field: 'name', sortable: true, order: "desc" },
        { title: "规格数量", field: 'productCnt', sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/product/template/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/product/template/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //产品规格
    displayTable("#table-product-product", "/product/product/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "规格编码", field: 'defaultCode', sortable: true, order: "desc" },
        { title: "规格类别", field: 'category', sortable: true, order: "desc" },
        { title: "产品规格", field: 'name', sortable: true, order: "desc" },
        { title: "产品款式", field: 'parent', sortable: true, order: "desc" },
        { title: "规格属性", field: 'attributes', align: "center", sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/product/product/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/product/product/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);

    //产品属性值
    displayTable("#table-product-attributevalue", "/product/attributevalue/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "属性", field: 'attribute', sortable: true, order: "desc" },
        { title: "属性值", field: 'value', align: "center", sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                html += "<a href='/product/attributevalue/" + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='/product/attributevalue/" + row.Id + "?action=detial' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);


});