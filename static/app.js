var App = {
  ws: null,

  openWebsocket: function(){
    if (App.ws) {
        return false;
    }
    App.ws = new WebSocket("ws://localhost:8080/ws");
    App.ws.onopen = function(evt) {
      $(".editor").val("websocket open");
    }
    App.ws.onclose = function(evt) {
        App.ws = null;
    }
    App.ws.onmessage = function(evt) {
      $(".editor").val($(".editor").val() + "\n" + evt.data);
    }
    App.ws.onerror = function(evt) {
        console.log(evt.data)
    }
  }
}

$(document).ready(function(){
  App.openWebsocket();
});
