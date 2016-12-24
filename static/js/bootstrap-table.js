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
        pageSize: 20,
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
            filter: true,
        }
        if (onExpandRow != undefined) {
            options.detailView = true;
            options.onExpandRow = onExpandRow;
        }
        $tableNode.bootstrapTable(options);
    };
    //用户表
    displayTable("#table-user", "/user/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "中文名称", field: 'namezh', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "部门", field: 'department', sortable: true, order: "desc", filter: { type: "select", data: [] } },
        { title: "职位", field: 'position', sortable: true, order: "desc", filter: { type: "select", data: [] } },
        { title: "邮箱", field: 'email', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "座机", field: 'tel', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "QQ", field: 'qq', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        { title: "微信", field: 'wechat', sortable: true, order: "desc", filter: { type: "input", data: [] } },
        {
            title: "管理员",
            field: 'isadmin',
            sortable: true,
            order: "desc",
            align: "center",
            filter: { type: "input", data: [] },
            formatter: function cellStyle(value, row, index) {
                var html = "";
                if (row.isadmin) {
                    html = '<i class="fa fa-check"></i>';
                } else {
                    html = '<i class="fa fa-remove"></i>';
                }
                return html;
            }
        },
        {
            title: "有效",
            field: 'active',
            sortable: true,
            order: "desc",
            align: "center",
            filter: { type: "input", data: [] },
            formatter: function cellStyle(value, row, index) {
                var html = "";
                if (row.active) {
                    html = '<i class="fa fa-check"></i>';
                } else {
                    html = '<i class="fa fa-remove"></i>';
                }
                return html;
            }
        },
        {
            title: "操作",
            align: "left",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                var url = "/user/";
                if (row.active) {
                    html += "<a href='" + url + row.Id + "?action=invalid' class='table-action btn btn-xs btn-danger'>无效&nbsp<i class='fa fa-close'></i></a>";
                } else {
                    html += "<a href='" + url + row.Id + "?action=active' class='table-action btn btn-xs btn-success'>有效&nbsp<i class='fa fa-check'></i></a>";
                }
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }

    ], function(index, row, $detail) {
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
    //权限
    displayTable("#table-group", "/group/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "权限组名", field: 'name', sortable: true, order: "desc" },
        {
            title: "有效",
            field: 'active',
            sortable: true,
            order: "desc",
            align: "center",
            formatter: function cellStyle(value, row, index) {
                var html = "";
                if (row.active) {
                    html = '<i class="fa fa-check"></i>';
                } else {
                    html = '<i class="fa fa-remove"></i>';
                }
                return html;
            }
        },
        { title: "定位", field: 'location', sortable: true, order: "desc" },
        { title: "描述", field: 'description', sortable: true },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {
                var html = "";
                var url = "/group/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }

    ]);
    //登录记录表
    displayTable("#table-record", "/record/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc" },
        { title: "邮箱", field: 'email', sortable: true, order: "desc" },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc" },
        { title: "开始时间", field: 'start_time', sortable: true, order: "desc" },
        {
            title: "结束时间",
            field: 'end_time',
            sortable: true,
            order: "desc",
            formatter: function cellStyle(value, row, index) {
                var html = "";
                console.log(row.end_time);
                if (row.end_time == "0001-01-01 00:00:00") {
                    html = "<small>非正常退出</small>";
                } else {
                    html = row.end_time;
                }

                return html;
            }
        },
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
                var url = "/address/country/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/address/province/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/address/city/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/address/district/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/product/attributevalue/";
                for (key in datas) {
                    html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + "</a>";
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
                var url = "/product/attribute/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/product/category/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/product/template/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/product/product/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
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
                var url = "/product/attributevalue/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);
    //产品单位类别

    displayTable("#table-product-uom-categ", "/product/uomcateg/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "计量单位类别", field: 'name', sortable: true, order: "desc" },
        {
            title: "计量单位",
            field: 'uoms',
            align: "center",
            formatter: function cellStyle(value, row, index) {
                var datas = row.uoms;
                var html = "";
                for (key in datas) {
                    html += "<a  class='display-block label label-success' href='/product/uom/" + key + "?action=detail'>" + datas[key] + "</a>";
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
                var url = "/product/uomcateg/";
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);
    //产品单位
    displayTable("#table-product-uom", "/product/uom/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "计量单位类别", field: 'category', sortable: true, order: "desc" },
        { title: "计量单位", field: 'name', align: "center", sortable: true, order: "desc" },
        { title: "类型", field: 'type', align: "center", sortable: true, order: "desc" },
        {
            title: "有效",
            field: 'active',
            align: "center",
            sortable: true,
            order: "desc",
            formatter: function cellStyle(value, row, index) {
                var html = "";
                if (row.active) {
                    html = '<i class="fa fa-check"></i>';
                } else {
                    html = '<i class="fa fa-remove"></i>';
                }
                return html;
            }
        },
        { title: "比率", field: 'factor', align: "center", sortable: true, order: "desc" },
        { title: "更大比率", field: 'factorInv', align: "center", sortable: true, order: "desc" },
        { title: "舍入精度", field: 'rounding', align: "center", sortable: true, order: "desc" },
        {
            title: "操作",
            align: "center",
            field: 'action',
            formatter: function cellStyle(value, row, index) {

                var html = "";
                var url = "/product/uom/";
                if (row.active) {
                    html += "<a href='" + url + row.Id + "?action=invalid' class='table-action btn btn-xs btn-danger'>无效&nbsp<i class='fa fa-close'></i></a>";
                } else {
                    html += "<a href='" + url + row.Id + "?action=active' class='table-action btn btn-xs btn-success'>有效&nbsp<i class='fa fa-check'></i></a>";
                }
                html += "<a href='" + url + row.Id + "?action=edit' class='table-action btn btn-xs btn-info'>编辑&nbsp<i class='fa fa-pencil'></i></a>";
                html += "<a href='" + url + row.Id + "?action=detail' class='table-action btn btn-xs btn-primary'>详情&nbsp<i class='fa fa-external-link'></i></a>";
                return html;
            }
        }
    ]);
    //---------------------------------------款式中的属性列表----------------------------------
    //产品款式中产品属性信息
    displayTable("#one-product-template-attribute", "/product/template/?action=attributes", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        {
            title: "属性名",
            field: 'name',
            formatter: function cellStyle(value, row, index) {
                console.log(row);
                var html = ' <select name="productAttributeID1" id="productAttributeID1"  class="form-control select-product-attribute"></select>';
                return html;
            }
        },
        {
            title: "属性值",
            field: 'childs',
            align: "center",
            formatter: function cellStyle(value, row, index) {
                var datas = row.values;
                var html = "";
                var url = "/product/attributevalue/";
                for (key in datas) {
                    html += "<a  class='display-block label label-success' href='" + url + key + "?action=detail'>" + datas[key] + "</a>";
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
                var url = "/product/template/";
                html += "<a href='" + url + row.Id + "?action=invalid' class='table-action btn btn-xs btn-danger'><i class='fa fa-trash-o'></i></a>";
                return html;
            }
        }
    ]);
    $(function() {
        $("#add-one-product-template-attribute").click(function(e) {
            function newData() {
                var newRow = [{
                    id: 0,
                    name: "",
                    childs: "",
                    action: ""
                }];
                return newRow;
            }
            $("#one-product-template-attribute").bootstrapTable('append', newData());
        });
    });
    //bootstrap filter的input样式改为bootstrap,该行代码必须在所有的bootstrapTable后面，否则不能修改样式
    $(".fht-cell input").addClass("form-control");
    // 搜索功能关闭
    $("#search-disable").on('click', function(e) {
        $("#search-enable").removeClass("search-hidden");
        $("#search-disable").addClass("search-hidden");
        $(".fht-cell").addClass("search-hidden");
    });
    // 搜索功能开启
    $("#search-enable").on('click', function(e) {
        $(".fht-cell").removeClass("search-hidden");
        $("#search-enable").addClass("search-hidden");
        $("#search-disable").removeClass("search-hidden");
    });
    //鼠标进入展示详情，离开隐藏详情
    $('#table-user').on('mouseenter mouseleave', 'tbody>tr',
        function(e) {
            $(this).find(".detail-icon").trigger("click");
        }
    );

});