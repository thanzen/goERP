$(document).ready(function() {
    $.extend($.fn.dataTable.defaults, {
        language: {
            sProcessing: "处理中...",
            sLengthMenu: "每页 _MENU_ 项",
            sZeroRecords: "没有匹配结果",
            sInfo: "第 _START_ 至 _END_ 项，共 _TOTAL_ 项",
            sInfoEmpty: "第 0 至 0 项，共 0 项",
            sInfoFiltered: "(由 _MAX_ 项结果过滤)",
            sInfoPostFix: "",
            sSearch: "搜索:",
            sUrl: "",
            sEmptyTable: "没有符合条件的数据",
            sLoadingRecords: "载入中...",
            sInfoThousands: ",",
            oPaginate: {
                sFirst: "首页",
                sPrevious: "上页",
                sNext: "下页",
                sLast: "末页",
                sJump: "跳转"
            },
            oAria: {
                sSortAscending: ": 以升序排列此列",
                sSortDescending: ": 以降序排列此列"
            }
        }
    });

    //登录记录表
    var tableRecord = $('#table-record').DataTable({
        "Dom": "<'row-fluid'<'span6'l><'span6'f>r>t<'row-fluid'<'span6'i><'span6'p>>",
        lengthMenu: [
            [20, 10, 80, 100, 200, 500, 1000, 5000],
            [20, 10, 80, 100, 200, 500, 1000, 5000]
        ],
        // language: tableLang,
        ordering: true,
        renderer: "bootstrap",
        autoWidth: true,
        processing: true,
        serverSide: true,
        stateSave: true, //保存状态
        deferRender: true, //延迟渲染
        pagingType: "full_numbers", //分页风格
        dom: "<'row'<'col-sm-12'T>>" +
            "<'row'<'col-md-6 pull-left'l><'col-md-6 pull-right'f>>" +
            "<'row'<'col-sm-12'tr>>" +
            "<'row'<'col-sm-5'i><'col-sm-7'p>>",
        tableTools: {
            "sSwfPath": "/static/plugins/DataTables-1.10.12/extensions/TableTools-2.2.4/swf/copy_csv_xls_pdf.swf"
        },
        columnDefs: [{
            targets: [1],
            orderData: [1, 5] //如果第一列进行排序，有相同数据则按照第二列顺序排列
        }],
        ajax: {
            "url": "/record/",
            "type": "POST",
            'dataType': 'json',
            "data": function(params) {
                var xsrf = $("input[name ='_xsrf']")
                if (xsrf != undefined) {
                    params._xsrf = xsrf[0].value;
                }
                var limit = $("select[name='table-record_length']");
                if (limit != undefined) {
                    limit = (limit[0] && limit[0].value) || 20;
                    params.length = limit;
                }
                return params
            },
            "dataSrc": function(response) {
                return response.data;
            },


        },
        columns: [{
                "class": 'details-control',
                "orderable": false,
                "data": null,
                "defaultContent": ''
            },
            { "data": "email" },
            { "data": "mobile" },
            { "data": "username" },
            { "data": "namezh" },
            { "data": "start_time" },
            { "data": "end_time" },
            { "data": "ip" },
        ]

    });

    //地区表
    var tableDistrict = $('#table-district').DataTable({
        ordering: true,
        renderer: "bootstrap",
        autoWidth: true,
        processing: true,
        serverSide: true,
        stateSave: true, //保存状态
        deferRender: true, //延迟渲染
        pagingType: "full_numbers", //分页风格
        lengthMenu: [
            [20, 10, 80, 100, 200, 500, 1000, 5000],
            [20, 10, 80, 100, 200, 500, 1000, 5000]
        ],
        ajax: {
            "url": "/district/",
            "type": "POST",
            'dataType': 'json',
            "data": function(params) {
                var xsrf = $("input[name ='_xsrf']")
                if (xsrf != undefined) {
                    params._xsrf = xsrf[0].value;
                }
                var limit = $("select[name='table-district_length']");
                if (limit != undefined) {

                    limit = (limit[0] && limit[0].value) || 20;
                    params.length = limit;
                }
                return params
            },
            "dataSrc": function(response) {
                return response.data;
            },

        },
        columns: [{
                "class": 'details-control',
                "orderable": false,
                "data": null,
                "defaultContent": ''
            },
            { "data": "name" },
            { "data": "city" },
            { "data": "province" },
        ]
    });
});