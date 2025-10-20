<template>
  <div class="Tracks section">
    <h2>Tracks</h2>
    <div class="container">
      <button v-for="(track, index) in tracks" :key="index" class="track-btn">
        <img :src="track.src" />
        <h3>{{ track.name }}</h3>
        <p>{{ track.desc }}</p>
      </button>
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
  gap: 32px;
}

.track-btn {
  border-radius: var(--border-radius);
  width: 100%;
  cursor: pointer;
  border: none;
}
</style>