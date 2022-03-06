// websocket management
const ws = new WebSocket("ws://" + window.location.host + "/ws");

ws.onopen = () => {
  console.log("ws opened");
  ws.send("joined");
};

ws.onclose = () => {
  console.log("ws closed");
  ws.send("closed");
};

ws.onmessage = (msg) => {
  if (msg.data) {
    console.log(msg.data);
  }
};
