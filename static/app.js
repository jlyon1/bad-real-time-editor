var App = {
  ws: null,
  count: 0,
  oldDocumentValue: "",

  openWebsocket: function(){
    if (App.ws) {
        return false;
    }
    App.ws = new WebSocket("ws://127.0.0.1:8080/ws");
    App.ws.onopen = function(evt) {

    }
    App.ws.onclose = function(evt) {
        App.ws = null;
    }
    App.ws.onmessage = function(evt) {
      $("message").html(evt.data);
    }
    App.ws.onerror = function(evt) {
        console.log(evt.data)
    }
    $(document).keyup(function(e){
        var data = {method: 'delta', body: $(".editor").val()};
        App.ws.send(JSON.stringify(data));
    });

  },

  createDelta: function(type, val){
    return

  }
}

$(document).ready(function(){
  App.openWebsocket();
  App.oldDocumentValue = $(".editor").val();

});
