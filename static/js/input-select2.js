$(function() {
    $.fn.select2.defaults.set("language", "zh-CN");

    var LIMIT = 5;
    var select2AjaxData = function(selectClass, ajaxUrl, placeholder) {
        $(selectClass).select2({
            placeholder: placeholder,
            ajax: {
                url: ajaxUrl,
                dataType: 'json',
                delay: 250,
                type: "POST",
                data: function(params, page) {

                    var xsrf = $("input[name ='_xsrf']")[0].value;
                    return {
                        name: params.term, // search term
                        offset: params.page || 0,
                        _xsrf: xsrf,
                        limit: LIMIT,

                    };
                },
                success: function(data) {
                    return data.data;
                },
                processResults: function(data, params) {
                    params.page = params.page || 0;
                    var result = {
                        results: data.data,
                        pagination: {
                            more: (params.page * 2) < data.total
                        }
                    };

                    return result;
                },
                cache: true
            },
            escapeMarkup: function(markup) { return markup; },
            // minimumInputLength: 1,
            dropdownCssClass: "bigdrop",
            templateResult: function(repo) {
                if (repo.loading) { return repo.text; }
                return repo.name;
            },
            templateSelection: function(repo) {
                return repo.name;
            }
        });
    };
    select2AjaxData(".select-department", "/department/?action=search", "选择部门");
    select2AjaxData(".select-position", "/position/?action=search", "选择职位");
    select2AjaxData(".select-group", "/group/?action=search", "选择分组");
    select2AjaxData(".select-product-category", "/product/category/?action=search", "选择产品类别");

});