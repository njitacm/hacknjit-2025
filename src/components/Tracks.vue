<template>
  <div class="Tracks section">
    <h2>Tracks</h2>
    <Carousel :items="tracks" :numVisible="1" :numScroll="1" :circular="true">
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
import Carousel from './Carousel.vue';
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
/* These styles are in the parent, not 'scoped' */
#app-container {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin: 60px auto;
  max-width: 900px;
}

/* This is the class we gave our items */
.item {
  height: 200px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  font-weight: bold;
  color: #555;
  border-radius: 8px;
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