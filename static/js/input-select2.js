$(function() {
    $(".select-department").select2({
        placeholder: "选择部门",
        language: "zh-CN",
        ajax: {
            url: "/department/search/",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params, page) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    page: params.page || 1,
                    _xsrf: xsrf,
                    offset: 3,
                };
            },
            processResults: function(data, params) {
                params.page = params.page || 1;
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
                    page: params.page || 1,
                    _xsrf: xsrf,
                    offset: 3,
                };
            },
            processResults: function(data, params) {
                params.page = params.page || 1;
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
            url: "/group/search/",
            dataType: 'json',
            delay: 250,
            type: "POST",
            data: function(params, page) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    page: params.page || 1,
                    _xsrf: xsrf,
                    offset: 3,
                };
            },
            processResults: function(data, params) {
                params.page = params.page || 1;
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

});