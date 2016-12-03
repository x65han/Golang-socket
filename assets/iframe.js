var playerControl = {
    videoId: "LGIcCaGvY2s",
    isReady: false
};
function onYouTubeIframeAPIReady() {
    new YT.Player('master-player', {
        height: '100%',
        width:  '100%',
        videoId: 'LGIcCaGvY2s',
        events: {
            'onReady': onPlayerReady,
            'onStateChange': onPlayerStateChange
        },
        playerVars:{
            modestbranding: 0,
            controls: 0,
            showinfo: 0,
            rel: 0
        }
    });
}
function onPlayerReady(e){playerControl.isReady = true;}
function onPlayerStateChange(e) {
    if(e.data==YT.PlayerState.ENDED){
        e.target.playVideo();
    }else if(e.data==YT.PlayerState.PLAYING){

    }else  if(e.data==YT.PlayerState.BUFFERING){

    }else if(e.data==YT.PlayerState.PAUSED){

    }
}
