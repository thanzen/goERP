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
window.onresize =function(){
    if(document.documentElement.clientWidth < 800){
        if ($(".table-action-dropdown")[0]==undefined){
            var action  = $(".table-action");
            
            for (var j=0;j<action.length;j++){
                var actionNode = action[j];
                var children = actionNode.children;
                var actionNodeHtml = '<div class="btn-group btn-sm table-action-dropdown">'
                                +'<button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown">'
                                +'操作 <span class="caret"></span>'
                                +' </button>'
                                +'<ul class="dropdown-menu table-action-dropdown-ul" role="menu">';
                for (var k=0;k<children.length;k++){
                    actionNodeHtml +=  "<li>"+children[k].innerHTML+"</li>";
                }
                actionNodeHtml +='</ul>'
                                +'</div>';
                actionNode.innerHTML = actionNodeHtml;
            }
        }
    }else{
        var tableDropdownUl = $(".table-action-dropdown-ul");
        for (var i=0;i<tableDropdownUl.length;i++){
            var children = tableDropdownUl[i].children;
            var html = "";
            for (var j=0;j<children.length;j++){
                html += '<span class="table-action-list">'+children[j].innerHTML+'</span>';
            }
            tableDropdownUl[i].parentNode.parentNode.innerHTML =html;
        }
    }
}