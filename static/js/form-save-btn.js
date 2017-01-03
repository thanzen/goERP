// form的change事件保存数据变化内容， 保存事件提交的内容为变化的内容
$(".post-from").on("change", function(e) {
    console.log(e);
});
// 保存事件处理
$(".form-save-btn").on("click", function(e) {
    console.log(e);
    // e.preventDefault();
});