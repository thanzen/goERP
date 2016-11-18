$(function() {


    $(".select-department").select2({
        language: "zh-CN",
        ajax: {
            url: "/department/search/",
            dataType: "json",
            delay: 250,
            type: "POST",

            data: function(params) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    _xsrf: xsrf,
                    page: params.page || 1,
                    offset: 5,
                };
            },
            processResults: function(data, params) {
                return {
                    results: data,
                }
            },
            cache: true,
        },
        escapeMarkup: function(markup) { return markup; },
        // minimumInputLength: 1,
        templateResult: function(repo) {
            return repo.name
        },
        templateSelection: function(repo) {
            return repo.name
        },
    });
    $(".select-postion").select2({
        language: "zh-CN",
        ajax: {
            url: "/position/search/",
            dataType: "json",
            delay: 250,
            type: "POST",

            data: function(params) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    _xsrf: xsrf,
                    page: params.page || 1,
                    offset: 1,
                };
            },
            processResults: function(data, params) {
                console.log(data);
                return {
                    results: data.items,
                    pagination: {
                        more: data.page * data.pageSize < data.total
                    }
                }
            },
            cache: true,
        },
        escapeMarkup: function(markup) { return markup; },
        // minimumInputLength: 1,
        templateResult: function(repo) {
            return repo.name
        },
        templateSelection: function(repo) {
            return repo.name
        },
    });
    $(".select-group").select2({
        language: "zh-CN",
        ajax: {
            url: "/group/search/",
            dataType: "json",
            delay: 250,
            type: "POST",

            data: function(params) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    _xsrf: xsrf,
                    page: params.page || 1,
                    offset: 5,
                };
            },
            processResults: function(data, params) {
                return {
                    results: data,
                }
            },
            cache: true,
        },
        escapeMarkup: function(markup) { return markup; },
        // minimumInputLength: 1,
        templateResult: function(repo) {
            return repo.name
        },
        templateSelection: function(repo) {
            return repo.name
        },
    });
});