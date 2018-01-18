var playerControl = {
    videoId: "KQ6zr6kCPj8",
    isReady: false,
    cpu: null,
};
function onYouTubeIframeAPIReady() {
    new YT.Player('master-player', {
        height: '100%',
        width: '100%',
        videoId: playerControl.videoId,
        events: {
            'onReady': onPlayerReady,
            'onStateChange': onPlayerStateChange
        },
        playerVars: {
            modestbranding: 0,
            controls: 0,
            showinfo: 0,
            rel: 0
        }
    });
}
function onPlayerReady(e) {
    playerControl.isReady = true;
    playerControl.cpu = e;
    console.log("Ready");
}
function onPlayerStateChange(e) {
    if (e.data == YT.PlayerState.ENDED) {
        console.log("ended");
    } else if (e.data == YT.PlayerState.PLAYING) {
        console.log("playing");
    } else if (e.data == YT.PlayerState.BUFFERING) {
        console.log("buffering");
    } else if (e.data == YT.PlayerState.PAUSED) {
        console.log("paused");
    } else if (e.data == YT.PlayerState.CUED) {
        console.log("cued");
    }
}
