<template>
   <div id="app">
    <div class="mt-5 mb-5">
        <div v-if="!isRecording">
      <button
        @click="getStream"
        v-show="canRecord"
        class="ml-10"
      >
        Start Recording 🎥</button
      >
</div>
      <div v-else>
        <button @click="stopStream"> Stop Screen Recording ❌ </button>
      </div>
  </div>
  <video class="center" height="500px" controls autoplay id="video"></video>
   </div>
</template>

<script>
export default {
 data() {
    return {
      isRecording:false,
      canRecord: true,
      // videoTrack: null,
      displayOptions: {
        video: {
          cursor: "always",
        },
        audio: {
          echoCancellation: true,
          noiseSuppression: true,
          sampleRate: 44100,
        },
      },
      videoData: []
    }
 },
  methods: {
   async getStream() {
       navigator.mediaDevices
        .getDisplayMedia(this.displayOptions)
        .then((stream) => {
          let video = document.getElementById("video");
          video.srcObject = stream;
         
          // this.videoTrack = stream.getVideoTracks()[0];

          let mediaRecorder = new MediaRecorder(stream);

          mediaRecorder.ondataavailable = function (e) {
                    console.log({e})
                    this.videoData.push(e.data)
                }

                mediaRecorder.onstop = () => {
                  this.downloadVideo()
 stream.getTracks().forEach(function (track) {
                        track.stop();
                    });
                }
        })

        .catch((e) => console.log(e));
    },
    downloadVideo(){
    
                   
    var blob = new Blob(this.videoData, { type: 'video/webm' });
    
    var url = window.URL.createObjectURL(blob);
    var a = document.createElement('a');
    a.style.display = 'none';
    a.href = url;
    a.download = 'test.webm';
    document.body.appendChild(a);
    a.click();
    this.clearData()
       

    },
    clearData(){
      this.videoData=[]
 document.body.removeChild(a);
        window.URL.revokeObjectURL(url);
    },
    stopStream(){}
  }
}
</script>

<style>

</style>