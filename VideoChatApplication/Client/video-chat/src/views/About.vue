<template>
  <div class="about">
    Hello
    <video autoPlay id="userVideo"></video>
    <video autoPlay id="partnerVideo"></video>
  </div>
</template>
<script>
export default {
  name: "About",
  data() {
    return {};
  },
  methods: {
    async openCamera() {
      const allDevices = await navigator.mediaDevices.enumerateDevices();
      const cameras = allDevices.filter(
        (device) => device.kind == "videoinput"
      );
      const constraints = {
        audio: true,
        video: {
          deviceId: cameras[0].deviceId,
        },
      };

      try {
        return await navigator.mediaDevices.getUserMedia(constraints);
      } catch (err) {
        console.log(err);
      }
    },
    useEffect() {
      this.openCamera().then((stream) => {
        const localVideo = document.getElementById("userVideo");
        localVideo.srcObject = stream;
      });
    },
  },
  created() {
    this.useEffect();
  },
};
