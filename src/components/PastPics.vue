<template>
  <div class="PastPics">
    <Galleria v-bind="slideshowProps" :value="images">
      <template #item="slotProps">
        <img :src="slotProps.item.src" :alt="`HackNJIT ${slotProps.item.year} photo`" style="width: 100%" />
      </template>
    </Galleria>
  </div>
</template>

<script>
import Galleria from 'primevue/galleria';
import images from '../data/past_pics';
import { getImageUrl } from '../util';

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
    const newPaths = JSON.parse(JSON.stringify(images));

    for (let i = 0; i < newPaths.length; i++) {
      newPaths[i].src = await getImageUrl(`past_pics/${newPaths[i].src}`);
    }

    this.images = newPaths;
  }
}
</script>

<style scoped>
.p-galleria {
  --p-galleria-border-width: 0px;
  --p-galleria-border-radius: var(--border-radius);
  --p-galleria-indicator-button-background: var(--hacknjit-primary);
  --p-galleria-indicator-button-hover-background: var(--hacknjit-secondary);
  --p-galleria-indicator-button-active-background: white;
  --p-galleria-nav-button-background: white;
  --p-galleria-nav-button-color: var(--hacknjit-primary);
  --p-galleria-nav-button-hover-background: var(--hacknjit-secondary);
  --p-galleria-nav-button-focus-ring-color: white;

  @media(pointer: coarse) {
    & {
      --p-galleria-indicator-button-hover-background: var(--p-galleria-indicator-button-background);
      --p-galleria-nav-button-hover-background: var(--p-galleria-nav-button-background);
      --p-galleria-nav-button-hover-color: var(--p-galleria-nav-button-color);
    }
  }
}


img {
  width: 100%;
  aspect-ratio: 16/9;
  object-fit: cover;
}
</style>
