window.onload = function(){
		// Key Stroke
		document.body.onkeydown = function(e) {
				var ev = e || event;var key = ev.keyCode;
				if(ev.keyCode == 13) {
						sendMessage();
				}else if(key == 16){
						socket.emit("status", "");
				}
		}
		// Establish TCP Connection
		establishConnection();
		$('#roomname').css("background-color", "crimson");
}
function establishConnection(){
		// Connect to Socket.io
		socket = io.connect("");
		// Socket Receiver
		socket.on('pong', function(data){
				 console.log("pong: " + data);
		});
		socket.on('pipe', function(data){
				 console.log("pipe: >>");
				 console.log(data);
		});
}
function sendMessage(){
		var message = document.getElementById("message-box").value.trim();
		socket.emit("ping", message);
		document.getElementById("message-box").value = "";
		console.log("ping: " + message);
}
function connectToRoom(){
		var roomName = document.getElementById("message-box").value.trim();
		socket.emit("connect to room", roomName)
		document.getElementById("message-box").value = "";
		console.log("Connect to: " + roomName);
}
function updateUsername(){
		var username = document.getElementById("message-box").value.trim();
		socket.emit("update username", username)
		document.getElementById("message-box").value = "";
		console.log("Update username: " + username);
}
