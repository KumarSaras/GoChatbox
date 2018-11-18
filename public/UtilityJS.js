function sendMessage() {
    var input = document.getElementById("message");
    this.sendMsg = input.value;
    //input.setAttribute("value", sendMsg);
    this.ws.send(
        JSON.stringify({
            email: this.email,
            username: this.username,
            message: this.sendMsg
        }
    ));
    this.sendMsg = '';

}

function createAccount() {

    this.email = document.getElementById("email").value;
    this.username = document.getElementById("username").value;
    
    this.ws = new WebSocket('ws://' + window.location.host + '/ws');
    this.ws.addEventListener('message', function(e) {
        var msg = JSON.parse(e.data);
        var recvMsg = document.getElementById("recvMsg");
        recvMsg.setAttribute("value", msg.message);
    });

}