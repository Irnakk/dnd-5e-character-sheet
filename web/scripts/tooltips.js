$(document).ready(function(){
    $('.hover').mouseenter(function () {
        let hint = $(this).attr('rel');
        $('#hint').css({ 'left': $(this).offset().left, 'top': $(this).offset().top - 50});
        $('#hint').fadeIn(200).text(hint);
    }).mouseout(function () {
        $('#hint').fadeOut( 100 );
    });
});