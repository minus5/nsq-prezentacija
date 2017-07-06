function NSQDemo(receiveHandler) {
	var ws;
	var wsUrl;
	var handler;

	this.Init = function (receiveHandler){
		this.wsUrl = "ws://localhost:8080/ws";
		this.handler = receiveHandler;
		this.InitWS();
	}

	/////////////////////////////////////////////////////////////
	// Commands

	this.SendServer = function(cmd) {
		if (null == this.ws) {
			this.InitWS();
			return;
		}
		switch(this.ws.readyState){
			case 0:
				console.log("WS Opening");
				break;
			case 2:
			case 3:
				console.log("WS Closing or closed");
				this.InitWS();
				break;
			default:
				this.ws.send(JSON.stringify(cmd));
				break;
		}
	}

	/////////////////////////////////////////////////////////////
	// WebSocket

	// Main thread WebSocket init
	this.InitWS = function(){
		this.ws = new WebSocket(this.wsUrl);
		this.ws.binaryType = 'arraybuffer';
		this.ws.onclose = this.wsClose.bind(this);
		this.ws.onerror = this.wsError.bind(this);
		this.ws.onmessage = this.wsMessage.bind(this);
	}

	// Handle main thread websocket message
	this.wsMessage = function(evt){
		if (this.handler) {
			this.handler(evt.data);
		}
	}

	this.wsClose = function(evt){
		console.log("WebSocket closed", this.wsUrl, evt);
	}

	this.wsError = function(evt){
		console.error("WebSocket error", this.wsUrl, evt);
	}

	// Default init on create
	this.Init(receiveHandler);
}
