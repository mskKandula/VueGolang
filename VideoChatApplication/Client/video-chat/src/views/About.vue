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
    return {
      userVideo: null,
      userStream: null,
      webSocketRef: null,
      peerRef: null,
      partnerVideo: null,
    };
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
         this.userVideo = stream;
        this.userStream = stream;
        const localVideo = document.getElementById("userVideo");
        localVideo.srcObject = stream;
        let url = new URL(
          `ws://10.4.0.30:8000/join?roomId=${this.$route.params.roomId}`
        );

        this.webSocketRef = new WebSocket(url.href);

        this.webSocketRef.onopen = () => {
          this.webSocketRef.send(JSON.stringify({ join: true }));
        };

        this.webSocketRef.onmessage = async (e) => {
          const message = JSON.parse(e.data);

          if (message.join) {
            this.callUser();
          }

          if (message.offer) {
            this.handleOffer(message.offer);
          }

      });
    },
      callUser() {
      console.log("Calling Other User");
     },
      handleOffer(offer) {
        console.log(offer)
      console.log("Received Offer, Creating Answer");
     }
  },
  created() {
    this.useEffect();
  },
};
