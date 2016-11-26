 var applyColresizable = function() {

     $("#table-record").colResizable({
         liveDrag: true,
         gripInnerHtml: "<div class='grip'></div>",
         draggingClass: "dragging",
         resizeMode: 'fit',

     });
 }
 applyColresizable();