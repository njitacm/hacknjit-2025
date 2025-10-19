<template>
  <div v-for="(track, index) in tracks" :key="index" class="Tracks container">
    <div>
      <img :src="track.src"/>
      <h1>{{ track.name }}</h1>
      <p>{{ track.desc }}</p>
    </div>
  </div>
</template>

<script>
import tracks from "../data/tracks";

export default {
  data() {
    return {
      tracks: tracks
    }
  }
}
</script>

<style scoped>
.container {
  display: grid;
}
</style>