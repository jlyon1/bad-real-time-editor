var App = {
  ws: null,
  count: 0,
  openWebsocket: function(){
    if (App.ws) {
        return false;
    }
    App.ws = new WebSocket("ws://127.0.0.1:8080/ws");
    App.ws.onopen = function(evt) {
      $(".editor").val("websocket open");
    }
    App.ws.onclose = function(evt) {
        App.ws = null;
    }
    App.ws.onmessage = function(evt) {
      $(".editor").val(evt.data);
    }
    App.ws.onerror = function(evt) {
        console.log(evt.data)
    }
    $(".editor").keyup(function(){
      App.count += 1;
      if (App.count == 2){
        App.count = 0;
        App.ws.send($(".editor").val());
      }
    });

  }
}

$(document).ready(function(){
  App.openWebsocket();
  var val = $(".editor").val();
  var oldVal = "";

});
