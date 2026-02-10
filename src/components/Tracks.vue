<template>
  <div id="Tracks" ref="sectionRef">
    <h2>Tracks</h2>
    <div class="tracks_section flex"> 
      <Card class="track_card" v-for="track in tracks" :key="tracks.name">
          <template #header>
            <img class="track_image" :alt="track.name" :src="track.imgSrc" />
          </template>
          <template #title>
            <div class="track_card_title"> {{ track.name }} </div>
          </template>
          <!-- <template #subtitle>{{ track.name }}</template> -->
          <template #content>
              <text-clamp :text='track.desc' :max-lines='4' class="text-clamp"/>
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
import TextClamp from 'vue3-text-clamp';

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

.text-clamp {
  font-size: 1em;
}

.tracks_section {
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
  justify-content: space-around;
  gap: 1em;
}

.track_card { 
  background-color: rgba(204, 255, 255, .4);
  width: 25rem; 
  overflow: hidden
}

.track_card_title { 
  font-size: 1.75em;
  text-align: center;
}

.track_desc {
  height: 6rem;
  overflow: hidden;
  text-overflow: ellipsis;
} 

.track_image {
  width: 80%;
  max-height: 275px;
  margin: 5% 10%;
}

.tracks_footer {
  display: flex; 
  justify-content: center; 
}
</style>
