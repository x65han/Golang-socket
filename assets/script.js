window.onload = function(){
	alert("script.js is working");
	socket = io.connect("");
	socket.emit("ping", "Johnson")
	socket.on('pong', function(data){
			 alert("server pong " + data);
	 });
}
