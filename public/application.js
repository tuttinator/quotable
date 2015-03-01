function getSelectedText() {
  var text = "";
  if (typeof window.getSelection != "undefined") {
    text = window.getSelection().toString();
  } else if (typeof document.selection != "undefined" && document.selection.type == "Text") {
    text = document.selection.createRange().text;
  }
  return text;
}

$(function() {
  $('.marketing').highlighter();
  $('body').on('click', 'button.twitter-share', function(e) {
    console.log('clicked');
    e.preventDefault();
    var selectedText = getSelectedText();
    var text = selectedText.match(/\b[\w']+(?:[^\w\n]+[\w']+){0,8}\b/g).join("\n")
    $('#twitter-share-modal').modal();
    $('.twitter-image').html('<span>' + selectedText + '</span>');
    var data = {
      url: window.location.href,
      text: text
    }

    $.ajax({
      type: 'POST',
      url: '/create',
      data: JSON.stringify(data),
      dataType: 'json',
      success: function(result) {
        $('.twitter-image').html('<img src="' + result.key + '.png" style="width: 100%">');
      }
    });
  });
});
