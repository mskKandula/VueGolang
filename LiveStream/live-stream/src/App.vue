<template>
  <div id="app">
    <LiveStreaming />
  </div>
</template>

<script>
import LiveStreaming from "./components/LiveStreaming.vue";

export default {
  name: "App",
  components: {
    LiveStreaming,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
