<template>
  <div id="Tracks" ref="sectionRef">
    <h2>Tracks</h2>
    <div class="tracks_section"> 
      <Card style="width: 25rem; overflow: hidden" v-for="track in tracks" :key="tracks.name">
          <template #header>
              <img alt="user header" src="../assets/logos/discord.svg" />
          </template>
          <template #title>{{ track.name }}</template>
          <template #subtitle>{{ track.name }}</template>
          <template #content>
              <p class="m-0">
                  {{ track.desc }}
              </p>
          </template>
          <template #footer>
              <div class="flex gap-4 mt-1">
                  <Button label="Cancel" severity="secondary" variant="outlined" class="w-full" />
                  <Button label="Save" class="w-full" />
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
  for (const track of tracks) {
    images.value.push('url("' + await getImageUrl(`tracks/${track.bkgSrc}`) + '")');
  }
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

</style>
