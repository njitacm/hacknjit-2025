<template>
  <div class="PastPics" ref="sectionRef" id="Past-Pics">
    <div class="page-side-padding container">
      <img src="../assets/PastPictures/hacknjit2024_1.jpg" />
      <img src="../assets/PastPictures/hacknjit2024_2.jpg" id="two" />
      <img src="../assets/PastPictures/hacknjit2024_3.jpg" />
      <img src="../assets/PastPictures/hacknjit2024_4.jpg" />
    </div>
  </div>
</template>

<script>
import { useIntersectionObserver } from '../composables/useIntersectionObserver';

const { observe, unobserve } = useIntersectionObserver();

export default {
  mounted() {
    observe(this.$refs.sectionRef);
  },
  beforeUnmount() {
    unobserve(this.$refs.sectionRef);
  }
}
</script>

<style scoped>
.container {
  --gap: 16px;
  gap: var(--gap);
  display: grid;
  grid-template-columns: repeat(2, calc(50% - var(--gap) / 2));
  place-content: center;
  /* max-width: var(--max-content-width); */

  @media(max-width: 600px) {
    & {
      grid-template-columns: 100%;
    }
  }
}

img {
  width: 100%;
  height: auto;
  border-radius: 0.5rem;
  display: block;
  aspect-ratio: 16/9;
  object-fit: cover;
}

#two {
  /* adjust the numbers to crop */
  object-position: 100% 40%;
}
</style>
