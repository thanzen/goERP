$.fn.editable.defaults.mode = 'inline';
//---------------------------------------款式中的属性列表----------------------------------
//bootstrapTable
$("#one-product-template-attribute").bootstrapTable({
    url: "/product/template",
    queryParams: function(params) {
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf != undefined) {
            params._xsrf = xsrf[0].value;
        }
        params.action = 'table';
        return params;
    },

    columns: [
        { title: "全选", field: 'id', checkbox: true, align: "center", valign: "middle" },
        { title: "属性名称", field: 'name', sortable: true, order: "desc" },
        { title: "属性值", field: 'attributes', sortable: true, order: "desc" },
    ],
});
//x-editable