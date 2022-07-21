/*
 * Editiable Table
 */

$(function() {
  $('#mainTable').editableTableWidget().numericInputExample().find('td:first').focus();
  $('#textAreaEditor').editableTableWidget({
    editor: $('<textarea>')
  });
  window.prettyPrint && prettyPrint();
});