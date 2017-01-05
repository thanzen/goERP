// form的change事件保存数据变化内容， 保存事件提交的内容为变化的内容
$(".post-from").on("change", function(e) {
    // console.log(e);
});
// 保存事件处理
$(".form-save-btn").on("click", function(e) {
    console.log(e);
    // e.preventDefault();
});
// 图片上传处理
$('#product-template-images').fileinput({
    language: 'zh',
    uploadUrl: '#',
    uploadExtraData: (function() {
        var params = {};
        var xsrf = $("input[name ='_xsrf']");
        if (xsrf.length > 0) {
            params._xsrf = xsrf[0].value;
        }
        params.action = "uploadFile";
        params._method = "PUT";
        return params;
    })(),
    allowedFileExtensions: ['jpg', 'png', 'gif'],
});
$(".form-disabled .file-input").hide();
$("#productTemplateForm .form-edit-btn").bind("click.images", function() {
    $(".file-input").show();
});
$("#productTemplateForm .form-save-btn,#productTemplateForm .form-cancel-btn").bind("click.images", function() {
    $(".file-input").hide();
});
// 单击图片悬浮
$(".click-modal-view").dblclick(function(e) {
    var imageSrc = e.currentTarget.src;
    $("#productImage").attr("src", imageSrc);
    $('#productImagesModal').modal('show');
});
// 款式form中图片懒加载
$('a[href="#productImages"]').on('shown.bs.tab', function(e) {
    // 图片加载
    $("#productImages .click-modal-view").each(function(index, el) {
        if ($(el).attr("src") == "") {
            $(el).attr("src", $(el)[0].dataset["src"]);
        }
    });

})