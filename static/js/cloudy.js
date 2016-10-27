$(function(){
    $(".list-page-info").change(function(){
        var page = $("#page-page");
        var offset = $("#page-offset");
        if(page && offset){
            page = page[0].value;
            offset = offset[0].value;
            location.replace("/user/list/?page="+page+"&offset="+offset) ;
             
        }
        
    });
});