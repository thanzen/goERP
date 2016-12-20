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
    $(".select-product-uom-category-type").on("change", function(e) {
        if (e.currentTarget.value == "1") {
            $("#factorInvDisplay").addClass("hidden");
            $("#factorDisplay").removeClass("hidden");
        } else if (e.currentTarget.value == "3") {
            $("#factorDisplay").addClass("hidden");
            $("#factorInvDisplay").removeClass("hidden");
        } else {
            $("#factorDisplay").addClass("hidden");
            $("#factorInvDisplay").addClass("hidden");
        }
    });
});