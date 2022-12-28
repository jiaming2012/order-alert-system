// api/index.js
var socket = new WebSocket("ws://localhost:8080/orders");

let connect = cb => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = wsEvent => {
        try {
            const orders = JSON.parse(wsEvent.data);
            cb(orders);
        } catch (err) {
            console.error(`failed to parse wsEvent: ${err}`);
            return;
        }
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };