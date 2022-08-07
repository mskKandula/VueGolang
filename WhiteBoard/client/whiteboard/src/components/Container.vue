<template>
  <div className="container">
    <div class="tools-section">
      <div className="color-picker-container">
        Select Brush Color : &nbsp;
        <input
          type="color"
          v-model="selectedColor"
          @change="changeColor()"
        />
      </div>

      <div className="brushsize-container">
        Select Brush Size : &nbsp;
        <select
          v-model="size"
          @change="changeSize()"
        >
          <option>5</option>
          <option>10</option>
          <option>15</option>
          <option>20</option>
          <option>25</option>
          <option>30</option>
        </select>
      </div>
    </div>

    <div class="board-container">
      <Board ></Board>
    </div>
  </div>
</template>

<script>
import Board from './Board.vue'

export default {
     components: {
    Board
  },
  data(){
    return{
         color: "#000000",
            size: "5"
    }
  },
  methods:{
changeColor(){

},
changeSize(){

}
  }
};
</script>

<style>
.container {
    position: fixed;
    width: 100%;
    height: 100%;
    background: linear-gradient(45deg, #03A9F4, #4CAF50);
}

.color-picker-container {
    text-align: center;
    display: inline;
    color: white;
}

.brushsize-container {
    display: inline;
    color: white;
    margin-left: 12px;
}

.tools-section {
    text-align: center;
}

.board-container {
    width: 90%;
    height: 90%;
    margin: auto;
    margin-top: 1%;
    background: white;
}
</style>