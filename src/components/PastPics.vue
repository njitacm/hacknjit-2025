<template>
  <div class="PastPics section" ref="sectionRef" id="Past-Pics">
    <h2 class="section-title">Past Pictures</h2>
    <Galleria v-bind="slideshowProps" :value="images" containerClass="slideshow" >
      <template #item="slotProps">
        <img :src="slotProps.item.src" :alt="`HackNJIT ${slotProps.item.year} photo`" style="width: 100%" />
      </template>
    </Galleria>
  </div>
</template>

<script>
import Galleria from 'primevue/galleria';
import images from '../data/past_pics';
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
import { getImageUrl } from '../util';
const { observe, unobserve } = useIntersectionObserver();

export default {
  components: { Galleria },
  data() {
    return {
      images: null,
      slideshowProps: {
        value: images,
        showItemNavigators: true,
        showItemNavigatorsOnHover: true,
        showThumbnails: false,
        showIndicators: true,
        circular: true,
        autoPlay: true,
        transitionInterval: 5000,
      },
    };
  },
  async mounted() {
    observe(this.$refs.sectionRef);
    const newPaths = JSON.parse(JSON.stringify(images));

    for (let i = 0; i < newPaths.length; i++) {
      newPaths[i].src = await getImageUrl(`past_pics/${newPaths[i].src}`);
    }

    this.images = newPaths;
    console.log(this.images);
  },
  beforeUnmount() {
    unobserve(this.$refs.sectionRef);
  }
}
</script>

<style scoped>
.slideshow {
  max-width: 1000px;
  margin: 0 auto;
}

img {
  width: 100%;
  aspect-ratio: 16/9;
  object-fit: cover;
}
</style>
