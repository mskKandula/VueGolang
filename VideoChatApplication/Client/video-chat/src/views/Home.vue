<template>
  <div class="home">
    <button @click="createRoom">Create Room</button>
  </div>
</template>

<script>
export default {
  name: "Home",
  methods: {
    async createRoom() {
      const resp = await fetch("http://localhost:8000/create");
      console.log(resp);
      const { roomId } = await resp.json();
      console.log("15", roomId);
      this.$router.push(`/room/${roomId}`);
    },
  },
};
</script>
