<template>
  <div class="sketch" id="sketch">
    <canvas className="board" id="board"></canvas>
  </div>
</template>

<script>
export default {
  props: {
    color: String,
    size: String,
  },
  data() {
    return {
      ctx: null,
      canvas: null,
      mouse: null,
      last_mouse:null,
      timeout:null
    };
  },
  methods: {
    drawOnCanvas() {
       this.canvas = document.querySelector("#board");
      this.ctx = this.canvas.getContext("2d");

      var sketch = document.querySelector("#sketch");
      var sketch_style = getComputedStyle(sketch);
      this.canvas.width = parseInt(sketch_style.getPropertyValue("width"));
      this.canvas.height = parseInt(sketch_style.getPropertyValue("height"));

       this.mouse = { x: 0, y: 0 };
       this.last_mouse = { x: 0, y: 0 };

      /* Mouse Capturing Work */
      this.canvas.addEventListener(
        "mousemove",
        function (e) {
          this.last_mouse.x = this.mouse.x;
          this.last_mouse.y = this.mouse.y;

          this.mouse.x = e.pageX - this.offsetLeft;
          this.mouse.y = e.pageY - this.offsetTop;
        },
        false
      );

      /* Drawing on Paint App */
      this.ctx.lineWidth = this.size;
      this.ctx.lineJoin = "round";
      this.ctx.lineCap = "round";
      this.ctx.strokeStyle = this.color;

      this.canvas.addEventListener(
        "mousedown",
        function () {
          this.canvas.addEventListener("mousemove", this.onPaint, false);
        },
        false
      );

      this.canvas.addEventListener(
        "mouseup",
        function () {
          this.canvas.removeEventListener("mousemove", this.onPaint, false);
        },
        false
      );
    },
    onPaint(){
 this.ctx.beginPath();
            this.ctx.moveTo(this.last_mouse.x, this.last_mouse.y);
            this.ctx.lineTo(this.mouse.x, this.mouse.y);
            this.ctx.closePath();
            this.ctx.stroke();

            if(this.timeout != undefined){ 
                clearTimeout(this.timeout)}

            this.timeout = setTimeout(function(){
                let base64ImageData = this.canvas.toDataURL("image/png");
               console.log(base64ImageData)
            })
  },
  }
};
</script>

<style>
.board {
  width: 100%;
  height: 100%;
}

.sketch {
  width: 100%;
  height: 100%;
}
</style>