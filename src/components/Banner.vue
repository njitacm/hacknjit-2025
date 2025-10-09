<template>
  <div class="Banner" id="HackNJIT" ref="sectionRef" :style="{ background: gradient }">
    <!-- <NotifSignup id="NotifSignup" /> -->
    <h1 class="title">HackNJIT</h1>
    <div class="earth-container">
      <img src="..\assets\globe.svg" class="earth">
    </div>
    <div class="spacer"></div>
    <div class="content">
      <RouterLink to="/registration" class="router-link pill register-button">Register Now</RouterLink>
      <TheCountdown />
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
  height: 100svh;
  min-height: 900px;
  display: grid;
  grid-template-rows: 1fr auto;
  padding-bottom: 32px;
}

.title {
  top: 250px;
  font-weight: bold;
  font-size: 10em;
  position: absolute;
  width: 100%;
  text-align: center;
  z-index: -2;
}

.earth-container {
  display: inline-block;
  width: 1000px;
  position: absolute;
  overflow: hidden;
  top: -100px;
  left: 50%;
  transform: translateX(-50%);
  z-index: -1;
}

.earth {
  object-fit: cover;
  -webkit-mask-image: -webkit-linear-gradient(to bottom, rgba(0, 0, 0, 1) 35%, rgba(0, 0, 0, 0) 75%);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 1) 35%, rgba(0, 0, 0, 0) 75%);
}

.content {
  display: grid;
  justify-content: center;
  justify-items: center;
  gap: 32px;
}

.register-button {
  font-size: 2em;
}

.view-hint {
  width: 100%;
  /* left: 50%; */
  /* transform: translateX(-50%); */
  /* bottom: 32px; */
  /* position: absolute; */
  align-content: center;
}

@media(max-width: 1000px) {
  .title {
    font-size: 6em;
    top: 310px;
  }

}

@media(max-width: 600px) {
  .title {
    font-size: 4em;
    top: 345px;
  }

  .register-button {
    font-size: 1.75em;
  }
}
</style>