<template>
  <div class="Tracks section">
    <h2>Tracks</h2>
    <Carousel :value="tracks" :numVisible="1" :numScroll="1" :circular="true" :showIndicators="false"
      containerClass="carousel-container" contentClass="carousel-container">
      <template #item="{ data, index }">
        <div class="track-container">
          <div class="track-sub-container">
            <div :style="{ backgroundImage: images[index] }" class="bkg-img"></div>
            <h3>{{ data.name }}</h3>
            <p>{{ data.desc }}</p>
          </div>
        </div>
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
.p-carousel-next-button {
  transform: scale(1.5) !important;
}

.carousel-container,
.p-carousel-viewport {
  border-radius: var(--border-radius);
}

.track-container {
  padding: 32px;

}

.track-sub-container {
  position: relative;
  border-radius: var(--border-radius);
  width: 100%;
  cursor: pointer;
  border: none;
  padding: 16px;
  display: grid;
  justify-content: center;
  background-color: #00000075;
}

.track-sub-container .bkg-img {
  border-radius: inherit;
  position: absolute;
  background-size: 25px;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}
</style>