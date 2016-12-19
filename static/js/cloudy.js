$(document).ready(function() {
    //form-readonly下所有的输入下所有的输入框readonly
    $(".form-readonly input.form-control").attr("readonly", "readonly");
    $(".form-readonly select").prop("disabled", true);
    $(".form-readonly .form-save-btn,.form-readonly .form-cancel-btn").hide();
    //编辑删除readonly属性，输入框变成可编辑状态
    $(".form-edit-btn").on("click", function(e) {
        e.preventDefault();
        $(".form-readonly .form-edit-btn").hide();
        $(".form-readonly .form-save-btn,.form-readonly .form-cancel-btn").show();
        $(".form-readonly input.form-control").removeAttr("readonly");
        $(".form-readonly select").prop("disabled", false);
    });
    $(".form-cancel-btn").on("click", function(e) {
        e.preventDefault();
        $(".form-readonly .form-edit-btn").show();
        $(".form-readonly .form-save-btn,.form-readonly .form-cancel-btn").hide();
        $(".form-readonly input.form-control").attr("readonly", "readonly");
        $(".form-readonly select").prop("disabled", true);
    });

});