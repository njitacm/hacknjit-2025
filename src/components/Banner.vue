<template>
  <div class="Banner" id="HackNJIT" ref="sectionRef" :style="{ background: gradient }">
    <!-- <NotifSignup id="NotifSignup" /> -->
    <div class="title-super-container">
      <div class="title-container">
        <h1 class="title">HackNJIT</h1>
        <img src="..\assets\HackNJIT2025\globe-half-of-the-earth.png" class="earth">
        <TheCountdown class="countdown" />
      </div>
    </div>
    <div class="view-hint">
      <RouterLink :to="{ hash: '#Past-Pics' }" class="router-link">View last year's photos</RouterLink>
      <br />
      <br />
      <RouterLink :to="{ hash: '#Past-Pics' }">
        <img src="../assets/down_arrow.svg" style="width: 25px; height: 25px" />
      </RouterLink>
    </div>
  </div>
</template>

<script>
import NotifSignup from "./NotifSignup.vue";
import TheCountdown from "./TheCountdown.vue";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
const { observe, unobserve } = useIntersectionObserver();

export default {
  props: {
    gradient: String
  },
  components: {
    NotifSignup,
    TheCountdown
  },
  mounted() {
    observe(this.$refs.sectionRef);
  },
  beforeUnmount() {
    unobserve(this.$refs.sectionRef);
  }
}
</script>

<style scoped>
.Banner {
  min-height: 100svh;
  height: fit-content;
  grid-template-rows: 1fr auto;
  gap: 1rem;
  position: relative;
  overflow-x: hidden;
}

.title-super-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.title-container {
  position: relative;
  top: 128px;
  /* transform: translateY(50%); */
  /* transform-origin: 50% 50%; */
}

.title {
  font-weight: bold;
  font-size: 10em;
  position: absolute;
  width: 100%;
  text-align: center;
}

.earth {
  width: 1000px;
  position: relative;
  transform-origin: 50% 50%;
  margin: 0;
}

.countdown {
  bottom: 32px;
  position: absolute;
}

.view-hint {
  left: 50%;
  transform: translateX(-50%);
  bottom: 32px;
  position: absolute;
  align-content: center;
}

.view-hint .router-link {
  color: white;
  text-decoration: none;
}

@media(max-width: 1000px) {
  .title {
    font-size: 6em;
    top: 60px;
  }
}

@media(max-width: 600px) {
  .title {
    font-size: 4em;
    top: 95px;
  }
}
</style>