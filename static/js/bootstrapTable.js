$(document).ready(function(){
  
    $.extend($.fn.bootstrapTable.defaults,{
        method:"post",
        dataType:"json",
        locale: "zh-CN",
        contentType:"application/x-www-form-urlencoded",
        sidePagination: "server",
        
        // onClickRow: function(row, $element) {
        //     //$element是当前tr的jquery对象
        //     $element.css("background-color", "green");
        // },//单击row事件
    });
    var $tableCountry =$("#table-country");
    $tableCountry.bootstrapTable({
        url:"/address/country/",
        queryParams:function(params){
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action= 'table';
            return params;
        },
        dataField:"data",
        pagination:true,
        pageNumber:1,
        pageSize:10,
        pageList:[10, 25, 50, 100],
        columns:[
            {title:"全选",field:'id',checkbox:true,align:"center",valign:"middle"},
            {title:"国家",field:'name',sortable:true,order:"desc"},
        ],
        
    });
    var $tableProvince =$("#table-province");
    $tableProvince.bootstrapTable({
        url:"/address/province/",
        queryParams:function(params){
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action= 'table';
            return params;
        },
        dataField:"data",
        pagination:true,
        pageNumber:1,
        pageSize:10,
        pageList:[10, 25, 50, 100],
        columns:[
            {title:"全选",field:'id',checkbox:true,align:"center",valign:"middle"},
            {title:"省份",field:'name',sortable:true,order:"desc"},
            {title:"国家",field:'country',sortable:true,order:"desc"},

        ],
        
    });
    var $tableCity =$("#table-city");
    $tableCity.bootstrapTable({
        url:"/address/city/",
        queryParams:function(params){
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action= 'table';
            return params;
        },
        dataField:"data",
        pagination:true,
        pageNumber:1,
        pageSize:10,
        pageList:[10, 25, 50, 100],
        columns:[
            {title:"全选",field:'id',checkbox:true,align:"center",valign:"middle"},
            {title:"城市",field:'name',sortable:true,order:"desc"},
            {title:"省份",field:'province',sortable:true,order:"desc"},
            {title:"国家",field:'country',sortable:true,order:"desc"},
            
        ],
        
    });
    var $tableDistrict =$("#table-district");
    $tableDistrict.bootstrapTable({
        url:"/address/district/",
        queryParams:function(params){
            var xsrf = $("input[name ='_xsrf']");
            if (xsrf != undefined) {
                params._xsrf = xsrf[0].value;
            }
            params.action= 'table';
            return params;
        },
        dataField:"data",
        pagination:true,
        pageNumber:1,
        pageSize:10,
        pageList:[10, 25, 50, 100],
        height:function () {
            return document.body.offsetHeight;
          },
        columns:[
            {title:"全选",field:'id',checkbox:true,align:"center",valign:"middle"},
            {title:"地区",field:'name',sortable:true,order:"desc"},
            {title:"城市",field:'city',sortable:true,order:"desc"},
            {title:"省份",field:'province',sortable:true,order:"desc"},
            {title:"国家",field:'country',sortable:true,order:"desc"},
            
        ],
        
    });
});
