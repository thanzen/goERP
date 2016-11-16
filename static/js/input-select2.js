$(".select-department").select2({
        language: "zh-CN", 
        ajax:{
            url:"/department/search/",
            dataType:"json",
            delay:250,
            type:"POST",
            
            data:function(params){
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    _xsrf:xsrf,
                };
            },
            processResults:function(data,params){
                 return {
                    results: data ,
                }
            },
            cache:true,
        },
        escapeMarkup: function (markup) { return markup; },
        minimumInputLength: 1,
        templateResult: function(repo){
            return repo.name
        }, 
        templateSelection: function(repo){
            return repo.name
        },  
    });
    $(".select-postion").select2({
        language: "zh-CN", 
        ajax:{
            url:"/department/search/",
            dataType:"json",
            delay:250,
            type:"POST",
            
            data:function(params){
                var xsrf = $("input[name ='_xsrf']")[0].value;
                return {
                    name: params.term, // search term
                    _xsrf:xsrf,
                };
            },
            processResults:function(data,params){
                 return {
                    results: data ,
                }
            },
            cache:true,
        },
        escapeMarkup: function (markup) { return markup; },
        minimumInputLength: 1,
        templateResult: function(repo){
            return repo.name
        }, 
        templateSelection: function(repo){
            return repo.name
        },  
    });