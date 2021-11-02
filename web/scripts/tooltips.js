$(document).ready(function(){
    $('.hover').mouseenter(function(e){
        let hint = $(this).attr('rel');
        $('#hint').css({ 'left': e.clientX - 25, 'top': e.clientY - 50});
        $('#hint').show().text(hint);
    }).mouseout(function () {
        $('#hint').hide();
    });
});