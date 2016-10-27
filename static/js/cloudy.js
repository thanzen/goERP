$(function(){
    $(".list-page-info").change(function(){
        var page = $("#page-page");
        var offset = $("#page-offset");
        if(page && offset){
            page = page[0].value;
            var url = offset[0].dataset.url;
            offset = offset[0].value;
            location.replace(url+"/list/?page="+page+"&offset="+offset) ;
        }
        
    });
    
	//tree视图下checkbox选择操作
	$(".checkbox-top").bind("click",function(e){
         var checked = e.currentTarget.checked;
		//显示顶部操作按钮
		if (checked){
			$(".list-top-action").css("display","block");
            $(".checkbox-data").attr("checked","checked");
		}else{
			$(".list-top-action").css("display","none");
            $(".checkbox-data").attr("checked","");
		}
		 
	});
	 
});