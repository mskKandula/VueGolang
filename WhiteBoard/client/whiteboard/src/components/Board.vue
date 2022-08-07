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
    };
  },
  methods: {
    drawOnCanvas() {
      var canvas = document.querySelector("#board");
      this.ctx = canvas.getContext("2d");

      var sketch = document.querySelector("#sketch");
      var sketch_style = getComputedStyle(sketch);
      canvas.width = parseInt(sketch_style.getPropertyValue("width"));
      canvas.height = parseInt(sketch_style.getPropertyValue("height"));

      var mouse = { x: 0, y: 0 };
      var last_mouse = { x: 0, y: 0 };

      /* Mouse Capturing Work */
      canvas.addEventListener(
        "mousemove",
        function (e) {
          last_mouse.x = mouse.x;
          last_mouse.y = mouse.y;

          mouse.x = e.pageX - this.offsetLeft;
          mouse.y = e.pageY - this.offsetTop;
        },
        false
      );

      /* Drawing on Paint App */
      this.ctx.lineWidth = this.size;
      this.ctx.lineJoin = "round";
      this.ctx.lineCap = "round";
      this.ctx.strokeStyle = this.color;

      canvas.addEventListener(
        "mousedown",
        function () {
          canvas.addEventListener("mousemove", this.onPaint, false);
        },
        false
      );

      canvas.addEventListener(
        "mouseup",
        function () {
          canvas.removeEventListener("mousemove", this.onPaint, false);
        },
        false
      );
    },
    onPaint(){

    }
  },
};
</script>

<style>
</style>