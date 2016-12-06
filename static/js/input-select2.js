$(function() {
    $.fn.select2.defaults.set("language", "zh-CN");

    var LIMIT = 5;
    $(".select-department").select2({
        placeholder: "选择部门",
        ajax: {
            url: "/department/search/",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params, page) {

                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    page: params.page || 0,
                    _xsrf: xsrf,
                    offset: LIMIT,
                };
            },
            success: function(data) {
                return data.items;
            },
            processResults: function(data, params) {
                params.page = params.page || 1;
                var result = {
                    results: data.items,
                    pagination: {
                        more: (params.page * 2) < data.total_count
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
    $(".select-position").select2({
        placeholder: "选择部门",
        language: "zh-CN",
        ajax: {
            url: "/position/search/",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params, page) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    page: params.page || 0,
                    _xsrf: xsrf,
                    offset: LIMIT,
                };
            },
            processResults: function(data, params) {
                params.page = params.page || 0;
                var result = {
                    results: data.items,
                    pagination: {
                        more: (data.page * data.pageSize) < data.total_count
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
    $(".select-group").select2({
        placeholder: "选择部门",
        language: "zh-CN",
        ajax: {
            url: "/group/",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params, page) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    start: params.page || 0,
                    _xsrf: xsrf,
                    length: LIMIT,
                };
            },
            processResults: function(data, params) {
                params.page = params.page || 1;
                var result = {
                    results: data.data,
                    pagination: {
                        more: data.page < data.pages
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
    $(".select-product-category").select2({
        placeholder: "选择产品类别",
        language: "zh-CN",
        ajax: {
            url: "/product/category/",
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
                params.page = params.page || 1;
                var result = {
                    results: data.data,
                    pagination: {
                        more: data.page < data.pages
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

});