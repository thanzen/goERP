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
    allowedFileExtensions: ['jpg', 'png', 'gif'],
    initialPreview: [
        "<img src='https://img.alicdn.com/tps/i4/TB1DK8cNpXXXXXCXXXXSutbFXXX.jpg_490x490Q80S0.jpg' class='file-preview-image' alt='Desert' title='Desert'>",
    ],

});
// 款式form中图片懒加载
$('a[href="#images"]').on('shown.bs.tab', function(e) {
    // 图片加载
    $("#images");
    console.log(e);
})