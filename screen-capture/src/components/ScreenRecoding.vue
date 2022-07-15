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
      videoTrack: null,
      displayOptions: {
        video: {
          cursor: "always",
        },
        audio: {
          echoCancellation: true,
          noiseSuppression: true,
          sampleRate: 44100,
        },
      }
    }
 },
  methods: {
   async getStream() {
       navigator.mediaDevices
        .getDisplayMedia(this.displayOptions)
        .then((stream) => {
          let video = document.getElementById("video");
          video.srcObject = stream;
         
          this.videoTrack = stream.getVideoTracks()[0];
        })

        .catch((e) => console.log(e));
    },
    stopStream(){}
  }
}
</script>

<style>

</style>