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
              <p class="track_desc">{{ track.desc }}</p>
          </template>
      </Card>
    </div>
  </div>
</template>

<script setup>
import Card from 'primevue/card';
import { onMounted, onUnmounted, useTemplateRef } from "vue";
import tracks from "../data/tracks";
import { getImageUrl } from "../util";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();
const sectionRef = useTemplateRef("sectionRef");

onMounted(async () => {
  observe(sectionRef.value);
});

onUnmounted(() => {
  unobserve(sectionRef.value);
});
</script>

<style scoped>

#Tracks {
  margin: 0% 5%;
}

.track_desc {
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
  width: 30%; 
  overflow: hidden
}

.track_card_title { 
  font-size: 1.75em;
  text-align: center;
}

.track_image {
  width: 80%;
  max-height: 275px;
  margin: 5% 10%;
}

@media (max-width: 1350px) { 
  .track_card_title { 
    font-size: 1.25em;
  }
  .track_desc {
    font-size: 0.9em;
  }
  .track_image {
    max-height: 250px;
  }
}


@media (max-width: 1250px) { 
  .track_image {
    max-height: 200px;
  }
}

@media (max-width: 1000px) { 
  .track_card_title { 
    font-size: 1em;
  }
  .track_image {
    max-height: 150px;
  }
}

@media (max-width: 850px) { 
  .track_desc {
    font-size: 0.9em;
  } 
}

@media (max-width: 800px) { 
  .track_card { 
    width: 75%; 
  }
  .track_card_title { 
    font-size: 1.5em;
  }
  .track_desc {
    font-size: 1.15em;
  } 
  .track_image {
    max-height: unset;
    max-width: 60%;
    margin: 1% 20%;
  }
}

@media (max-width: 500px) { 
  .track_desc {
    font-size: 1.1em;
  } 
}

</style>
