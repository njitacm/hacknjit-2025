<template>
  <div id="Tracks" ref="sectionRef">
    <h2>Tracks</h2>
    <div class="tracks_section flex"> 
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
                  <Button label="Learn More" class="w-full" @click="createModal(track)"/>
              </div>
          </template>
      </Card>
    </div>
    <Dialog v-model:visible="modal_visible" modal :header="modal_title" :style="{ width: '50rem' }" :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">
      <p class="mb-8">
            {{ modal_p1 }}
      </p>
    </Dialog>
  </div>
</template>

<script setup>
import Card from 'primevue/card';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { onMounted, onUnmounted, ref, useTemplateRef } from "vue";
import tracks from "../data/tracks";
import { getImageUrl } from "../util";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();
const sectionRef = useTemplateRef("sectionRef");

const modal_visible = ref(false);
const modal_title = ref("");
const modal_p1 = ref("");

onMounted(async () => {
  observe(sectionRef.value);
});

onUnmounted(() => {
  unobserve(sectionRef.value);
});

function createModal(track) {
  modal_visible.value = true;
  modal_title.value = track.name;
  modal_p1.value = track.desc;
}
</script>

<style scoped>

#Tracks {
  margin: 0% 5%;
}

.tracks_section {
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
  justify-content: space-around;
  gap: 1em;
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
