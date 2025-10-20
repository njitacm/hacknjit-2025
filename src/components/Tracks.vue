<template>
  <div class="Tracks section">
    <h2>Tracks</h2>
    <div class="container">
      <button v-for="(track, index) in tracks" :key="index" class="track-btn"
        :style="{ backgroundImage: images[index] }">
        <h3>{{ track.name }}</h3>
        <p>{{ track.desc }}</p>
      </button>
    </div>
  </div>
</template>

<script>
import tracks from "../data/tracks";
import { getImageUrl } from "../util";

export default {
  data() {
    return {
      tracks: tracks,
      images: []
    }
  },
  async mounted() {
    for (const track of tracks) {
      this.images.push('url("' + await getImageUrl(`tracks/${track.bkgSrc}`) + '")');
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