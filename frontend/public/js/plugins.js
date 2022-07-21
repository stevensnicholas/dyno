/*Preloader*/
$(window).on('load', function () {
  setTimeout(function () {
    $('body').addClass('loaded');
  }, 200);
});