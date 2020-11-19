export default function newSocket(callback) {
    let socket = new WebSocket("ws://localhost:8080/ws")

    socket.onmessage = function(e) {
        callback(JSON.parse(e.data))
    }

    socket.onclose = function(e) {
        if (e.wasClean) {
            alert(`[close] Connection closed cleanly, code=${e.code} reason=${e.reason}`)
        } else {
            alert('[close] Connection died')
        }
    }

    socket.onerror = function(e) {
        alert(`[error] ${e.message}`)
    }

    return socket
}
