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
      this.peerRef = await this.createPeer();

      this.userStream.getTracks().forEach((track) => {
        this.peerRef.addTrack(track, this.userStream);
      });
     },
      handleOffer(offer) {
        console.log("Received Offer, Creating Answer");
      this.peerRef = await this.createPeer();

      await this.peerRef.setRemoteDescription(
        new RTCSessionDescription(offer)
      );

      this.userStream.getTracks().forEach((track) => {
        this.peerRef.addTrack(track, this.userStream);
      });

      const answer = await this.peerRef.createAnswer();
      await this.peerRef.setLocalDescription(answer);

      this.webSocketRef.send(
        JSON.stringify({ answer: this.peerRef.localDescription })
      );
     },
      createPeer() {
       return new Promise((resolve) => {
      console.log("Creating Peer Connection");
       const peer = new RTCPeerConnection({
        iceServers: [
          {
            urls: [
              "stun:stun.l.google.com:19302",
              "stun:stun2.l.google.com:19302",
            ],
          },
        ],
      });
      peer.onnegotiationneeded = this.handleNegotiationNeeded;
      peer.onicecandidate = this.handleIceCandidateEvent;
      peer.ontrack = this.handleTrackEvent;
      resolve(peer)})
    },
    async handleNegotiationNeeded() {
      console.log("Creating Offer",);

      try {
        const myOffer = await this.peerRef.createOffer();

         await this.peerRef.setLocalDescription(myOffer);
console.log("139",this.peerRef.localDescription)
        this.webSocketRef.send(
          JSON.stringify({ offer: this.peerRef.localDescription })
        );
      } catch (err) {
        console.log("136", err);
      }
    },
  },
  created() {
    this.useEffect();
  },
};
