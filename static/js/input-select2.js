$(function() {
    $(".select-department").select2({
        placeholder: "选择部门",
        language:"zh-CN",
        ajax: {
            url: "/department/search/",
            dataType: 'json',
            delay: 250,
            type:"POST",
            data: function (params) {
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    page: params.page ||1,
                    _xsrf:xsrf,
                    offset:1,
                };
            },
            processResults: function (data, params) {
                console.log(data);
                params.page = params.page || 1;
                var result = {
                    results: data.items,
                    pagination: {
                        more: (params.page * data.pageSize) < data.total_count
                    }
                };
                return result;
            },
            cache: true
        },
        escapeMarkup: function (markup) { return markup; },  
        templateResult: function(repo){
            if (repo.loading) return repo.text;
            return  repo.name
        },  
        templateSelection: function(repo){
            return repo.name  
        }  
        });
});