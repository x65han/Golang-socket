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
		$('#roomname').css("color", "white");
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
		socket.on('play', function(data){
				if(playerControl.cpu != null){
						playerControl.cpu.target.playVideo();
						var duration = playerControl.cpu.target.getDuration();
						console.log("Total time: " + duration);
				}else{
						alert("player control cpu error");
				}
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
function seekTo(time){
	if(playerControl.cpu != null){
			playerControl.cpu.target.seekTo(time);
	}else{
			alert("player control cpu error");
	}
}
function requestToPlay(){
		socket.emit("request to play", playerControl.videoId);
}
function requestToStop(){
		if(playerControl.cpu != null){
				playerControl.cpu.target.pauseVideo()
		}else{
				alert("player control cpu error");
		}
}
