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
    var displayTable = function(selectId, ajaxUrl, columns) {
            var $tableNode = $(selectId);
            $tableNode.bootstrapTable({
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
            });
        }
        //用户表
    displayTable("#table-user", "/user/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc" },
        { title: "中文名称", field: 'namezh', sortable: true, order: "desc" },
        { title: "部门", field: 'department', sortable: true, order: "desc" },
        { title: "邮箱", field: 'email', sortable: true, order: "desc" },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc" },
        { title: "座机", field: 'tel', sortable: true, order: "desc" },
        { title: "QQ", field: 'qq', sortable: true, order: "desc" },
        { title: "微信", field: 'wechat', sortable: true, order: "desc" },
        { title: "管理员", field: 'isadmin', sortable: true, order: "desc" },
        { title: "有效", field: 'active', sortable: true, order: "desc" },

    ]);
    //登录记录表
    displayTable("#table-record", "/record/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "用户名", field: 'username', sortable: true, order: "desc" },
        { title: "邮箱", field: 'email', sortable: true, order: "desc" },
        { title: "手机号码", field: 'mobile', sortable: true, order: "desc" },
        { title: "开始时间", field: 'start_time', sortable: true, order: "desc" },
        { title: "结束时间", field: 'end_time', sortable: true, order: "desc" },
        { title: "IP地址", field: 'ip', sortable: true, order: "desc" },

    ]);
    //国家表
    displayTable("#table-country", "/country/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "国家", field: 'name', sortable: true, order: "desc" },
    ]);
    //省份表
    displayTable("#table-province", "/province/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "省份", field: 'name', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
    ]);

    //城市表
    displayTable("#table-city", "/city/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "城市", field: 'name', sortable: true, order: "desc" },
        { title: "省份", field: 'province', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
    ]);
    //区县表
    displayTable("#table-district", "/district/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "地区", field: 'name', sortable: true, order: "desc" },
        { title: "城市", field: 'city', sortable: true, order: "desc" },
        { title: "省份", field: 'province', sortable: true, order: "desc" },
        { title: "国家", field: 'country', sortable: true, order: "desc" },
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
            formatter: function cellStyle(col, row, d) {
                var datas = row.values;
                var html = "";
                for (key in datas) {
                    html += "<span  class='attribute-value-table-display label label-primary'>" + datas[key] + "</span>";
                }
                return html;
            }
        },
    ]);

    //产品类别
    displayTable("#table-product-category", "/product/category/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "类别名", field: 'name', sortable: true, order: "desc" },
        { title: "上级", field: 'parent', sortable: true, order: "desc" },
        { title: "上级路径", field: 'path', sortable: true, order: "desc" },
    ]);

    //产品款式
    displayTable("#table-product-template", "/product/template/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "款式编码", field: 'defaultCode', sortable: true, order: "desc" },
        { title: "款式类别", field: 'category', sortable: true, order: "desc" },
        { title: "产品款式", field: 'name', sortable: true, order: "desc" },
        { title: "规格数量", field: 'productCnt', sortable: true, order: "desc" },
    ]);

    //产品规格
    displayTable("#table-product-product", "/product/product/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "规格编码", field: 'defaultCode', sortable: true, order: "desc" },
        { title: "规格类别", field: 'category', sortable: true, order: "desc" },
        { title: "产品规格", field: 'name', sortable: true, order: "desc" },
        { title: "产品款式", field: 'parent', sortable: true, order: "desc" },
        { title: "规格属性", field: 'attributes', sortable: true, order: "desc" },
    ]);

    //产品属性值
    displayTable("#table-product-attributevalue", "/product/attributevalue/", [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "属性", field: 'attribute', sortable: true, order: "desc" },
        { title: "属性值", field: 'value', sortable: true, order: "desc" },
    ]);


});