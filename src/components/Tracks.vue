<template>
  <div id="Tracks" ref="sectionRef">
    <h2>Tracks</h2>
    <div class="tracks_section"> 
      <Card style="width: 25rem; overflow: hidden" v-for="track in tracks" :key="tracks.name">
          <template #header>
            <img class="card_image" :alt="track.name" :src="track.imgSrc" />
          </template>
          <template #title>{{ track.name }}</template>
          <template #subtitle>{{ track.name }}</template>
          <template #content>
              <p class="m-0">
                  {{ track.desc }}
              </p>
          </template>
          <template #footer>
              <div class="tracks_footer">
                  <Button label="Learn More" class="w-full" />
              </div>
          </template>
      </Card>
    </div>
  </div>
</template>

<script setup>
import Card from 'primevue/card';
import Button from 'primevue/button';
import { onMounted, onUnmounted, ref, useTemplateRef } from "vue";
import tracks from "../data/tracks";
import { getImageUrl } from "../util";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();
const sectionRef = useTemplateRef("sectionRef");

// refs
const images = ref([]);

onMounted(async () => {
  observe(sectionRef.value);
});

onUnmounted(() => {
  unobserve(sectionRef.value);
});
</script>

<style scoped>

.tracks_section {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  margin: 0% 5%;
}

.card_image {
  width: 80%;
  margin: 10%;
}

.tracks_footer {
  display: flex; 
  justify-content: center; 
}
</style>
