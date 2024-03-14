window.addEventListener("load", function (evt) {
  var output = document.getElementById("output");
  var input = document.getElementById("input");
  var ws;

  var print = function (message) {
    var d = document.createElement("div");
    d.textContent = message;
    output.appendChild(d);
    output.scroll(0, output.scrollHeight);
  };

  document.getElementById("open").onclick = function (evt) {
    evt.preventDefault()
    if (ws) {
      return false;
    }
    ws = new WebSocket("ws://localhost:8000/ws");
    ws.onopen = function (evt) {
      print("OPEN");
    };
    ws.onclose = function (evt) {
      print("CLOSE");
      ws = null;
    };
    ws.onmessage = function (evt) {
      print("RESPONSE: " + evt.data);
    };
    ws.onerror = function (evt) {
      print("ERROR: " + evt.data);
    };
    return false;
  };

  document.getElementById("send").onclick = function (evt) {
    evt.preventDefault()
    if (!ws) {
      return false;
    }
    print("SEND: " + input.value);
    ws.send(input.value);
    return false;
  };

  document.getElementById("close").onclick = function (evt) {
    evt.preventDefault()
    if (!ws) {
      return false;
    }
    ws.close();
    return false;
  };
});
