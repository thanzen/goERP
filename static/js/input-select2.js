$(function() {
    $.fn.select2.defaults.set("language", "zh-CN");

    var LIMIT = 5;
    //现根据class选择，再根据ID绑定时间，用于后期一个页面多个相同select的情况
    var select2AjaxData = function(selectClass, ajaxUrl, placeholder) {
        $(selectClass).each(function(index, el) {
            if (el.id != undefined && el.id != "") {
                var $selectNode = $("#" + el.id);
                $selectNode.select2({
                    placeholder: placeholder,
                    //初始化数据
                    initSelection: function(element, callback) {
                        var node = $("#" + el.id);
                        var id = node.data("default-id");
                        var name = node.data("default-name");
                        callback({ id: id, name: name });
                    },
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

                    templateResult: function(repo) {
                        if (repo.loading) { return repo.text; }
                        return repo.name;
                    },
                    templateSelection: function(repo) {
                        return repo.name;
                    }
                });

            }
        });

    };
    select2AjaxData(".select-department", "/department/?action=search", "选择部门");
    select2AjaxData(".select-position", "/position/?action=search", "选择职位");
    select2AjaxData(".select-group", "/group/?action=search", "选择分组");
    select2AjaxData(".select-product-category", "/product/category/?action=search", "选择产品类别");
});