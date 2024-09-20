<template>
  <div class="outer-container" ref="container">
    <img
      v-for="i in 20"
      :key="i"
      v-intersection-observer="[onIntersectionObserver]"
      class="gearSpin"
      :class="[i % 2 == 0 ? 'gearSpinInverted' : '']"
      src="../assets/HackNJIT2024/gears/gear4.svg"
    />
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    if (target.classList.contains("gearSpinInverted")) {
      target.classList.add("fade-in-reverse");
    } else {
      target.classList.add("fade-in");
    }
  }
}
</script>

<style scoped>
.outer-container {
  display: flex;
  justify-content: center;
  margin: 0 auto;
  width: 67.5%;
  padding: 0;
  border: 4px var(--color4) solid;
  border-radius: 30px;
  padding: 0.25rem;
}
.gearSpin {
  width: 5%;
  animation: gearSpin 1s linear infinite normal none;
  filter: brightness(0) saturate(100%) invert(43%) sepia(9%) saturate(1781%)
    hue-rotate(342deg) brightness(89%) contrast(91%);
}
.gearSpinInverted {
  animation-direction: reverse;
}
@keyframes gearSpin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
  /* 100% {
    transform: rotate(360);
  } */
}
.fade-in {
  animation: gearSpin 1s linear infinite normal none,
    fadeIn 0.5s linear 1 normal forwards;
}
.fade-in-reverse {
  animation: gearSpin 1s linear infinite reverse none,
    fadeIn 0.5s linear 1 normal forwards;
}
@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>