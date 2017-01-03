$(document).ready(function() {
    $('.input-radio').iCheck({
        checkboxClass: 'icheckbox_square-green',
        radioClass: 'iradio_square-green',
        increaseArea: '20%'
    });
    //有checked的radio默认选中
    $("input.checked").iCheck("check");

    //form-disabled下所有的输入下所有的输入框disabled
    $(".form-disabled input.form-control[type='text']").attr("disabled", "disabled");
    $(".form-disabled select").prop("disabled", true);
    $(".form-disabled .form-save-btn,.form-disabled .form-cancel-btn").hide();
    $(".form-disabled .input-radio").iCheck("disable");
    $(".form-disabled textarea.form-control").attr("disabled", "disabled");
    $(".form-disabled .form-control").hide();
    $(".form-disabled span.select2").hide();
    //编辑删除readonly属性，输入框变成可编辑状态
    $(".form-edit-btn").on("click", function(e) {
        e.preventDefault();
        $(".form-disabled .form-edit-btn").hide();
        $(".form-disabled .form-save-btn,.form-disabled .form-cancel-btn").show();
        $(".form-disabled input.form-control").removeAttr("disabled");
        $(".form-disabled select").prop("disabled", false);
        $(".input-radio").iCheck("enable");
        $(".form-disabled textarea.form-control").removeAttr("disabled");
        $(".form-disabled .p-form-control").hide();
        $(".form-disabled .form-control").show();
        $(".form-disabled span.select2").show();
    });
    $(".form-cancel-btn").on("click", function(e) {
        e.preventDefault();
        $(".form-disabled .form-edit-btn").show();
        $(".form-disabled .form-save-btn,.form-disabled .form-cancel-btn").hide();
        $(".form-disabled input.form-control").attr("disabled", "disabled");
        $(".form-disabled select").prop("disabled", true);
        $(".input-radio").iCheck("disable");
        $(".form-disabled textarea.form-control").attr("disabled", "disabled");
        $(".form-disabled .p-form-control").show();
        $(".form-disabled .form-control").hide();
        $(".form-disabled span.select2").hide();
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

    //如果搜索添加不为空，增加提示样式
    $("#listViewSearch input").change(function(e) {
        e.currentTarget.value = e.currentTarget.value.trim();
        nums = $.grep($("#listViewSearch input"), function(el, index) {
            if (el.value != "") {
                return true
            } else {
                return false
            }
        });
        if (nums.length > 0) {
            if ($("button[id^='clearListSearchCond']:first").hasClass("hide")) {
                $("button[id^='clearListSearchCond']").toggleClass("hide");
            }
        } else {
            if (!$("button[id^='clearListSearchCond']:first").hasClass("hide")) {
                $("button[id^='clearListSearchCond']").toggleClass("hide");
            }
        }
    });
    // 若过滤条件不为空， 显示清空条件按钮
    (function() {
        nums = $.grep($("#listViewSearch input"), function(el, index) {
            if (el.value != "") {
                return true
            } else {
                return false
            }
        });
        if (nums.length < 1) {
            $("button[id^='clearListSearchCond']").toggleClass("hide");
        }
    })();
    $("button[id^='clearListSearchCond']").click(function(e) {
        $("#listViewSearch input").each(function() {
            this.value = "";
        });
        $(this).addClass("hide");
    });
});