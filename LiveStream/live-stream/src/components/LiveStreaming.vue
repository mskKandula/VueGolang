<template>
  <div id="app">
    <div class="mt-5 mb-5">
      <div v-if="!isRecording">
        <button @click="getStream" v-show="canRecord" class="ml-10">
          Start Streaming 🎥
        </button>
      </div>
      <div v-else>
        <button @click="stopStream">Stop Streaming ❌</button>
      </div>
    </div>
    <canvas id="canvas" hidden></canvas>
    <video class="center" height="500px" controls autoplay id="video"></video>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isRecording: false,
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
      videoData: [],
      videoTrack: null,
      ws: null,
    };
  },
  methods: {
    async getStream() {
      this.isRecording = true;
      let webS = this.ws;
      navigator.mediaDevices
        .getUserMedia(this.displayOptions)
        .then((stream) => {
          let video = document.getElementById("video");
          video.srcObject = stream;

          this.videoTrack = stream.getVideoTracks()[0];

          let mediaRecorder = new MediaRecorder(stream);

          mediaRecorder.start(0);

          mediaRecorder.ondataavailable = function (e) {
            // this.videoData.push(e.data);
            // e.data.text().then((txt) => {
            //   webS.send(txt);
            // });
            webS.send(e.data);
          };

          mediaRecorder.onstop = () => {
            stream.getTracks().forEach(function (track) {
              track.stop();
            });
          };
        })

        .catch((e) => console.log(e));
    }
};
</script>

<style>
</style>