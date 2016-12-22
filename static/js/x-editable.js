$(document).ready(function() {
    $.extend($.fn.editable.defaults, {
        mode: 'inline',
        anim: true,
    });
    $('#productAttributeID1').editable({
        select2: {
            placeholder: 'Select Country',
            allowClear: true,
            minimumInputLength: 3,
            id: function(item) {
                return item.CountryId;
            },
            ajax: {
                url: '/product/attribute/?action=search',
                dataType: 'json',
                data: function(term, page) {
                    console.log(123);
                    return { query: term };
                },
                results: function(data, page) {
                    return { results: data };
                }
            },
            formatResult: function(item) {
                return item.CountryName;
            },
            formatSelection: function(item) {
                return item.CountryName;
            },
            initSelection: function(element, callback) {
                return $.get('/getCountryById', { query: element.val() }, function(data) {
                    callback(data);
                });
            }
        }
    });
});