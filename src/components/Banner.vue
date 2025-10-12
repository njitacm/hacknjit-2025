<template>
  <div class="Banner" id="HackNJIT" ref="sectionRef" :style="{ background: gradient }">
    <!-- <h1 class="title">HackNJIT</h1> -->
    <!-- <CurvedText class="title" :radius="500" :arc="100" /> -->
    <!-- <CurveText class="curved-title">HackNJIT</CurveText> -->
    <div class="curved-text-container">
      <CurvedText text="HackNJIT" :radius="radius" :arc="arc" view-box-size-x="500"
        view-box-size-y="500" center-y="calc(100% - 20px)" :rotation="rotation" :debug="true" :center-offset-x="0"
        :center-offset-y="23" />
    </div>
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
  data() {
    return {
      rotation: -10,
      interval: null,
      radius: 245,
      arc: 90,
    };
  },
  created() {
    // this.mqList1000 = window.matchMedia(`(max-width: 1000px)`);
  },
  mounted() {
    observe(this.$refs.sectionRef);
    // this.mqList1000.addEventListener("change", this.onMediaQuery1000Change);
    // this.interval = () => {
    //   if (this.rotation < 8) {
    //     this.rotation += 0.1;
    //   } else {
    //     clearInterval(this.interval);
    //   }
    // };
    setInterval(this.interval, 5);
  },
  beforeUnmount() {
    // this.mqList1000.removeEventListener("change", this.onMediaQuery1000Change);
    unobserve(this.$refs.sectionRef);
  },
  methods: {
    onMediaQuery1000Change(e) {
      if (e.matches) {
        console.log("below 1000px");
        this.radius = 375;
        this.arc = 50;
      } else {
        this.radius = 200;
        this.arc = 80;
      }
    }
  }
}
</script>

<style scoped>
.Banner {
  /* from the base of top: 250px for .title (arbritrary value that was previously used); 
  tweaking --top-offset will position title and globe for all media queries accordingly */
  --top-offset: -100px;
  --title-font-size: 5em;
  --offset-x: 15px;
  min-height: 900px;
  display: grid;
  grid-template-rows: 1fr auto;
  padding-bottom: 32px;
  overflow: hidden;
}

.curved-text-container {
  position: absolute;
  left: 50%;
  transform: translateX(calc(-50% + var(--offset-x)));
  top: calc(-65px + var(--top-offset));
  height: 1000px;
  width: 1000px;
  max-width: 100vw;
  z-index: -20;
  overflow: hidden;
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

@media(pointer: coarse) {
  .Banner {
    --offset-x: 7px;
  }
}

@media(max-width: 1050px) {
  .Banner {
    --title-font-size: 4em;
  }
}

@media(max-width: 600px) {
  .register-button {
    font-size: 1.75em;
  }
}
</style>