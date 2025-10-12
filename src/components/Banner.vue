<template>
  <div class="Banner" id="HackNJIT" ref="sectionRef">
    <CurvedText class="curved-title" text="HackNJIT" :radius="245" :arc="90" :view-box-size-x="500" :view-box-size-y="500"
    center-y="calc(100% - 20px)" :debug="false" :center-offset-x="0" :center-offset-y="23.5"
    text-class="curved-title-text" svg-class="curved-title-svg" />
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
import TheCountdown from "./TheCountdown.vue";
import CurvedText from "./CurvedText.vue";
import { useIntersectionObserver } from '../composables/useIntersectionObserver';
const { observe, unobserve } = useIntersectionObserver();

export default {
  components: { TheCountdown, CurvedText },
  mounted() {
    observe(this.$refs.sectionRef);
    setInterval(this.interval, 5);
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
  /* small offset to line up the curved text with the Earth */
  --offset-x: 15px;
  min-height: 900px;
  display: grid;
  grid-template-rows: 1fr auto;
  padding-bottom: 32px;
  overflow: hidden;
}

.curved-title {
  position: absolute;
  left: 50%;
  transform: translateX(calc(-50% + var(--offset-x)));
  top: calc(-65px + var(--top-offset));
  height: 1000px;
  width: 1000px;
  max-width: calc(100vw - 2 * var(--offset-x));
  z-index: -20;
  overflow: hidden;
}

:deep(.curved-title-svg) {
  transform-origin: bottom;
  animation: rotate-in 1s ease forwards;
}

:deep(.curved-title-text) {
  font-size: 5em;
  font-weight: bold;
  fill: #ffffff;
  font-family: monospace;
  text-shadow: 0 0 5px #00aaff, 0 0 10px #00aaff;
  opacity: 0;
  animation: fade-in 2s linear forwards;
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
  -webkit-mask-image: -webkit-linear-gradient(to bottom, rgba(0, 0, 0, 1) 45%, rgba(0, 0, 0, 0) 60%);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 1) 45%, rgba(0, 0, 0, 0) 60%);
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

@keyframes fade-in {
  from {
    opacity: 0;
  }

  to {
    opacity: 1;
  }
}

@keyframes rotate-in {
  from {
    transform: rotate(-10deg);
  }

  to {
    transform: none;
  }
}

@media(pointer: coarse) {
  .Banner {
    --offset-x: 7px;
  }
}

@media(max-width: 1050px) {
  .Banner {
    --offset-x: 7px;
  }

  :deep(.curved-title-text) {
    font-size: 4em;
  }
}

@media(max-width: 600px) {
  :deep(.curved-title-text) {
    font-size: 3em;
  }

  .register-button {
    font-size: 1.75em;
  }
}

@media(max-width: 400px) {
  :deep(.curved-title-text) {
    font-size: 2.5em;
  }
}
</style>