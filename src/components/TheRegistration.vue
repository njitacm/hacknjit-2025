<template>
  <div ref="container">
    <GearDivider />
    <h1 v-intersection-observer="[onIntersectionObserver]">
      Registration Coming Soon
    </h1>
    <GearDivider />
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    console.log(target);
    target.classList.add("fade-in");
  }
}
</script>
<script>
import GearDivider from "./GearDivider.vue";

export default {
  components: {
    GearDivider,
  },
};
</script>


<style scoped>
.fade-in {
  animation-name: fade-in;
  animation-duration: 0.5s;
  /* animation-delay: 1.5s; */
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
div {
  margin: 4rem 0;
}
h1 {
  font-size: 5rem;
  margin: 1rem 0;
}
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s linear;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
@media (max-width: 1200px) {
  h1 {
    font-size: 3.5rem;
    width: 70%;
    margin: 1rem auto;
  }
}
@media (max-width: 500px) {
  h1 {
    font-size: 3rem;
  }
}
@media (max-width: 450px) {
  h1 {
    width: 80%;
    font-size: 2.75rem;
  }
}
</style>