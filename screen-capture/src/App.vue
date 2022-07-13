<template>
  <div id="app">
    <ScreenRecoding />
  </div>
</template>

<script>
import ScreenRecoding from './components/ScreenRecoding.vue'

export default {
  name: 'App',
  components: {
    ScreenRecoding
  }
}
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
