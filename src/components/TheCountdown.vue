<template>
  <div ref="gradient" class="gradient"></div>
  <div class="outer-container" ref="container">
    <Transition>
      <header v-show="containerIsVisible">
        <h1>{{ format(days) }}</h1>
        <GearColon class="gear-colon" />
        <h1>{{ format(hours) }}</h1>
        <GearColon />
        <h1>{{ format(minutes) }}</h1>
        <GearColon />
        <h1>{{ format(seconds) }}</h1>
      </header></Transition
    >
    <h3>
      {{ format(tweened_month.toFixed(0)) }} /
      {{ format(tweened_day.toFixed(0)) }} /
      {{ tweened_year.toFixed(0) }}
    </h3>
    <div class="hatguy-container">
      <HatGuy />
    </div>
  </div>

  <div class="countdown-gradient"></div>
</template>

<script setup>
import gsap from "gsap";
import { ref } from "vue";
import { useIntersectionObserver } from "@vueuse/core";

const container = ref(null);
const containerIsVisible = ref(false);

useIntersectionObserver(container, ([{ isIntersecting }], _) => {
  containerIsVisible.value = isIntersecting;
});
</script>

<script>
import GearColon from "./GearColon.vue";
import HatGuy from "./HatGuy.vue";
export default {
  data() {
    return {
      milliseconds: 0,
      seconds: 0,
      minutes: 0,
      ten_minutes: 0,
      hours: 0,
      ten_hours: 0,
      days: 0,
      ten_days: 0,
      date_month: null,
      date_day: null,
      date_year: null,
      tweened_month: 0,
      tweened_day: 0,
      tweened_year: 2000,
    };
  },
  watch: {
    date_year(n) {
      gsap.to(this, { duration: 5, tweened_year: Number(n) || 0 });
    },
    date_day(n) {
      gsap.to(this, { duration: 5, tweened_day: Number(n) || 0 });
    },
    date_month(n) {
      gsap.to(this, { duration: 5, tweened_month: Number(n) || 0 });
    },
  },
  components: {
    GearColon,
    HatGuy,
  },
  methods: {
    setTime() {
      var timeLeft = new Date("2024/11/02 10:00:00") - new Date();

      var daysLeft = timeLeft / 8.64e7;
      var hoursLeft = (timeLeft / 3.6e6) % 24;
      var minLeft = (timeLeft / 60000) % 60;
      var secsLeft = (timeLeft / 1000) % 60;
      var millisecondsLeft = (secsLeft % 1) * 1000;

      this.milliseconds = millisecondsLeft;
      this.seconds = secsLeft;
      this.minutes = minLeft;
      this.hours = hoursLeft;
      this.days = daysLeft;
    },
    format(number) {
      var s = "" + Math.floor(number);
      return s.padStart(2, "0");
    },
  },
  computed: {
    secondRotation() {
      //   var rotationPercent = 360 - (this.milliseconds / 1000) * 360;
      //   return "transform: rotate(" + rotationPercent + "deg);";
      return "";
    },
    minuteRotation() {
      var rotationPercent = 360 - (this.seconds / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
    hourRotation() {
      var rotationPercent = 360 - (this.minutes / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
    dayRotation() {
      var rotationPercent = 360 - (this.hours / 60) * 360;
      return "transform: rotate(" + rotationPercent + "deg);";
    },
  },
  mounted() {
    setInterval(() => {
      this.setTime();
    }, 1000);
    setTimeout(() => {
      this.date_day = 2;
      this.date_month = 11;
      this.date_year = 2024;
    }, 2000);
  },
};
</script>
<style scoped>
* {
  transition: all 0.1s linear;
  box-sizing: border-box;
}
.outer-container {
  color: white;
  height: fit-content;
  width: 75%;
  height: calc(100vh);
  margin: 0 auto;
  padding-top: 3.5rem;
  width: fit-content;
  position: relative;
  overflow: hidden;
}
.hatguy-container {
  position: absolute;
  bottom: -40px;
  left: 50%;
  transform: translateX(-50%);
  animation-name: fade-in;
  opacity: 0;
  animation-delay: 4s;
  animation-duration: 2.5s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
@media (max-height: 630px) {
  .hatguy-container {
    left: 90%;
  }
}
.gradient {
  background: url("../assets/HackNJIT2024/repeating-bg.webp");
  animation-name: fade-in;
  animation-duration: 10s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  width: 100%;
  height: 100vh;
  position: absolute;
  z-index: -10;
  top: 0;
}
header {
  margin: 0 auto;
  display: flex;
  gap: 0.5rem;
  align-content: center;
}
.gear-colon {
  align-self: center;
}
h1 {
  font-size: 12rem;
}
h3 {
  font-size: 6rem;
  animation-name: fade-in;
  opacity: 0;
  animation-delay: 2s;
  animation-duration: 2.5s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
.text-container {
  align-self: center;
  height: fit-content;
}
.text-container h1 {
  font-size: 10rem;
  border: 4px white solid;
  border-radius: 50px;
  padding: 1rem 1rem;
  width: fit-content;
  margin: 0 auto;
}

.countdown-gradient {
  background: linear-gradient(
    to right,
    #00000090,
    #00000060 10%,
    #00000035 20%,
    #00000035 80%,
    #00000060 90%,
    #00000090
  );
  width: 100vw;
  height: 100vh;
  position: absolute;
  left: 0;
  z-index: -10;
  top: 0;
}
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
.v-enter-active,
.v-leave-active {
  transition: opacity 5.5s linear;
  transition-delay: 1s;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
@media (max-width: 1500px) {
  h1 {
    font-size: 8.75rem;
  }
  .outer-container {
    height: calc(85vh);
  }
  .gradient {
    height: calc(85vh);
  }
  .countdown-gradient {
    height: calc(85vh);
  }
}
@media (max-width: 1500px) {
  header {
    margin-top: 77px;
  }
  h3 {
    margin-top: -30px;
  }
}
@media (max-width: 1000px) {
  h1 {
    font-size: 6.75rem;
  }
  h3 {
    font-size: 5rem;
  }
}
@media (max-width: 750px) {
  h1 {
    font-size: 5.25rem;
  }
  h3 {
    margin-top: 0px;
    font-size: 4.5rem;
  }
}
@media (max-width: 600px) {
  h1 {
    font-size: 4.5rem;
    z-index: 100;
  }
  h3 {
    margin-top: 0px;
    font-size: 4rem;
    z-index: 100;
  }
  .outer-container {
    height: calc(75vh);
  }
  .gradient {
    height: calc(75vh);
  }
  .countdown-gradient {
    height: calc(75vh);
  }
}
@media (max-width: 500px) {
  h1 {
    font-size: 4rem;
  }
  h3 {
    font-size: 3rem;
  }
}
@media (max-width: 450px) {
  h1 {
    font-size: 3.5rem;
  }
  h3 {
    font-size: 2.5rem;
  }
}
</style>