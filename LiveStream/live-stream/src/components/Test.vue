<template>
  <div>
    <video ref="videoPlayer" autoplay></video>
    <button @click="startStreaming">Start Streaming</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      videoStream: null,
    };
  },
  mounted() {
    navigator.mediaDevices
      .getUserMedia({ video: true })
      .then((stream) => {
        this.$refs.videoPlayer.srcObject = stream;
        this.videoStream = stream;
      })
      .catch((err) => {
        console.error("Unable to access webcam", err);
      });
  },
  beforeDestroy() {
    if (this.videoStream) {
      this.videoStream.getTracks().forEach((track) => {
        track.stop();
      });
    }
  },
  methods: {
    startStreaming() {
      const videoTrack = this.videoStream.getVideoTracks()[0];
      const command = `ffmpeg -f v4l2 -input_format mjpeg -video_size ${
        videoTrack.getSettings().width
      }x${
        videoTrack.getSettings().height
      } -i video="Integrated Camera" -c:v libx264 -pix_fmt yuv420p -preset ultrafast -f flv rtmp://localhost:9000/test`;
      // Send the command to the server using Ajax, WebSockets, or another appropriate method
      fetch("/api/stream", {
        method: "POST",
        body: JSON.stringify({ command }),
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Failed to start streaming");
          }
          console.log("Streaming started successfully");
        })
        .catch((err) => {
          console.error("Failed to start streaming", err);
        });
    },
  },
};
</script>
