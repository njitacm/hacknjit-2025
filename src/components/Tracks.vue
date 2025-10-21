<template>
  <div class="Tracks section">
    <h2>Tracks</h2>
    <Carousel :value="tracks" :numVisible="1" :numScroll="1" :circular="true" orientation="vertical" :showIndicators="false">
      <template #item="{data, index}">
        <button class="track-btn" :style="{ backgroundImage: images[index] }">
          <h3>{{ data.name }}</h3>
          <p>{{ data.desc }}</p>
        </button>
      </template>
    </Carousel>
  </div>
</template>

<script setup>
import Carousel from 'primevue/carousel';
import { onMounted, ref } from "vue";
import tracks from "../data/tracks";
import { getImageUrl } from "../util";

// refs
const images = ref([]);

onMounted(async () => {
  for (const track of tracks) {
    images.value.push('url("' + await getImageUrl(`tracks/${track.bkgSrc}`) + '")');
  }
});
</script>

<style scoped>
.container {
  display: grid;
  gap: 32px;
}

.track-btn {
  position: relative;
  border-radius: var(--border-radius);
  background-size: 25px;
  width: 100%;
  cursor: pointer;
  border: none;
  padding: 16px;
}

/* semi-transparent overlay */
.track-btn::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 0;
  border-radius: var(--border-radius);
}
</style>