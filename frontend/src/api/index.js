// api/index.js

let connect = cb => {
    const wsScheme = (window.location.protocol === "http:") ? "ws" : "wss";
    let socket = new WebSocket(`${wsScheme}://${window.location.host}/orders`);

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

        setTimeout(() => {
            connect(cb);
        }, 8000);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

export { connect };