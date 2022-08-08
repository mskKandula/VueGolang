<template>
  <div id="bg">
    <canvas
      ref="canvas"
      @mousemove="draw"
      @mousedown="beginDrawing"
      @mouseup="stopDrawing"
      @mouseleave="cancelDrawing"
      :width="canvasWidth"
      :height="canvasHeight"
    />
    <div id="bar">
      <div class="bar-item">
        <router-link to="/">home</router-link>|
        <a href="https://github.com/mskKandula/VueGolang/tree/main/WhiteBoard"
          >github</a
        >|
        <button @click="clearCanvas">clear</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Room",
  data() {
    return {
      socketConn: null,
      x: 0,
      y: 0,
      canvasWidth: 0,
      canvasHeight: 0,
      isDrawing: false,
    };
  },
  methods: {
    drawLine(x1, y1, x2, y2) {
      let ctx = this.$refs.canvas.getContext("2d");
      ctx.beginPath();
      ctx.strokeStyle = "black";
      ctx.lineWidth = 7;
      ctx.moveTo(x1, y1);
      ctx.lineTo(x2, y2);
      ctx.stroke();
      ctx.closePath();
    },
    draw(e) {
      if (this.isDrawing) {
        this.drawLine(this.x, this.y, e.offsetX, e.offsetY);
        this.x = e.offsetX;
        this.y = e.offsetY;
      }
    },
    beginDrawing(e) {
      this.x = e.offsetX;
      this.y = e.offsetY;
      this.isDrawing = true;
    },
    stopDrawing(e) {
      if (this.isDrawing) {
        this.drawLine(this.x, this.y, e.offsetX, e.offsetY);
        this.isDrawing = false;
        this.socketConn.send(
          JSON.stringify({
            type: 1,
            body: this.$refs.canvas.toDataURL("image/png"),
            id: "drawing",
          })
        );
      }
    },
    cancelDrawing() {
      this.isDrawing = false;
    },
    drawUpdate(url) {
      let image = new Image();
      let ctx = this.$refs.canvas.getContext("2d");
      image.onload = () => {
        ctx.drawImage(image, 0, 0);
      };
      image.src = url;
    },
    handleResize() {
      let state = this.$refs.canvas.toDataURL("image/png");
      this.canvasWidth = window.innerWidth;
      this.canvasHeight = window.innerHeight - 25;
      this.drawUpdate(state);
    },

    clearCanvas() {
      let ctx = this.$refs.canvas.getContext("2d");
      ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);

      this.socketConn.send(
        JSON.stringify({
          type: 1,
          body: this.$refs.canvas.toDataURL("image/png"),
          id: "drawing",
        })
      );
    },
  },
  mounted() {
    const url = new URL("ws://localhost:8082/ws");

    this.socketConn = new WebSocket(url.href);

    this.socketConn.onconnect = (evt) => {
      console.log("ws connected", evt);
    };

    this.socketConn.onmessage = (evt) => {
      let data = evt.data;

      data = data.split(/\r?\n/);

      console.log("104", JSON.parse(data[0]).body);

      this.drawUpdate(JSON.parse(data[0]).body);
      data = "";
    };

    this.handleResize();
    window.addEventListener("resize", this.handleResize);
  },
};
</script>

<style scoped>
* {
  padding: 0;
  margin: 0;
}
#bg {
  background-color: #111;
  height: 100vh;
  width: 100vw;
}
#bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white;
  height: 25px;
}
.bar-item {
  align-items: center;
  padding-left: 5px;
  padding-right: 5px;
}
button,
a {
  background: none;
  text-decoration: none;
  color: inherit;
  border: none;
  padding: 0;
  padding-right: 5px;
  font: inherit;
  cursor: pointer;
  outline: inherit;
}
canvas {
  background-color: white;
  display: block;
}
</style>