<template>
  <div class="Banner" id="HackNJIT" ref="sectionRef" :style="{ background: gradient }">
    <!-- <h1 class="title">HackNJIT</h1> -->
    <CurvedText />
    <div class="earth-container">
      <img src="..\assets\globe.svg" class="earth">
    </div>
    <!-- <canvas id="canv"></canvas> -->
    <div class="spacer"></div>
    <div class="content">
      <RouterLink to="/registration" class="router-link pill register-button">Register Now</RouterLink>
      <TheCountdown />
    </div>
  </div>
</template>

<script>
import TheCountdown from "./TheCountdown.vue";
import CurvedText from "./CurvedText.vue";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
const { observe, unobserve } = useIntersectionObserver();

export default {
  props: {
    gradient: String
  },
  components: {
    TheCountdown, CurvedText
  },
  mounted() {
    observe(this.$refs.sectionRef);
  },
  beforeUnmount() {
    unobserve(this.$refs.sectionRef);
  },
}
</script>

<style scoped>
.Banner {
  /* from the base of top: 250px for .title (arbritrary value that was previously used); 
  tweaking --top-offset will position title and globe for all media queries accordingly */
  --top-offset: -100px;
  min-height: 900px;
  display: grid;
  grid-template-rows: 1fr auto;
  padding-bottom: 32px;
  overflow: hidden;
}

#canv {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  top: calc(250px + var(--top-offset));
  width: 600px;
  height: 250px;
}

.title {
  top: calc(250px + var(--top-offset));
  font-weight: bold;
  font-size: 10em;
  position: absolute;
  width: 100%;
  text-align: center;
  z-index: -2;
  animation: title-anim 1s ease forwards;
}

.earth-container {
  max-width: 100vw;
  position: relative;
  top: calc(-100px + var(--top-offset));
  z-index: -1;
  left: 50vw;
  transform: translateX(-50%);
}

.earth {
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  left: 50%;
  transform: translateX(-50%);
  position: absolute;
  width: 1000px;
  object-fit: cover;
  -webkit-mask-image: -webkit-linear-gradient(to bottom, rgba(0, 0, 0, 1) 35%, rgba(0, 0, 0, 0) 60%);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 1) 35%, rgba(0, 0, 0, 0) 60%);
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

@keyframes title-anim {
  from {
    transform: translateY(25px);
    opacity: 0;
  }

  to {
    transform: none;
    opacity: 1;
  }
}

@media(max-width: 1000px) {
  .title {
    font-size: 6em;
    top: calc(310px + var(--top-offset));
  }

}

@media(max-width: 600px) {
  .title {
    font-size: 4em;
    top: calc(345px + var(--top-offset));
  }

  .register-button {
    font-size: 1.75em;
  }
}
</style>