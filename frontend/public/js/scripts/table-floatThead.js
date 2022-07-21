/*
 * Float Thead Table
 */

$(function() {
  $('table.demo1').floatThead({
    position: 'fixed',
    scrollingTop: 65
  });
  $('table.demo1 thead').css('background', '#ECEFF1');
});