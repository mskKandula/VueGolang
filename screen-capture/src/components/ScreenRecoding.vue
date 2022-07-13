<template>
   <div id="app">
    <div class="mt-5 mb-5">
        <div v-if="!isRecording">
      <button
        @click="getStream"
        v-show="canRecord"
        class="ml-10"
      >
        Start Recording 🎥</button
      >
</div>
      <div v-else>
        <button @click="stopStream"> Stop Screen Recording ❌ </button>
      </div>
  </div>
   </div>
</template>

<script>
export default {
 data() {
    return {
      isRecording:false,
      canRecord: true,
    }
 },
  methods: {
    getStream() {
    },
    stopStream(){}
  }
}
</script>

<style>

</style>