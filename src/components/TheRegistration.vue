<template>
  <div ref="container" class="container">
    <h1 v-intersection-observer="[onIntersectionObserver]">
      Registration Is Open!
    </h1>
    <p><a href="https://forms.gle/y9YzP376uPTueiqK7">Sign up now!</a></p>
    <!-- <GearDivider /> -->
    <img
      id="gear5"
      class="floating-gear gear-fade-in"
      src="../assets/HackNJIT2024/gears/gear5.svg"
    />
    <img
      id="gear6"
      class="floating-gear gear-fade-in"
      src="../assets/HackNJIT2024/gears/gear6.svg"
    />
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    // console.log(target);
    target.classList.add("fade-in");
  }
}
</script>
<script>
export default {
  components: {},
};
</script>

<style scoped>
#gear6 {
  right: 15%;
  top: 0;
  width: 25vw;
  max-width: 782px;
}
#gear5 {
  left: 0%;
  width: 47vw;
  max-width: 782px;
}
.floating-gear {
  position: absolute;
  z-index: -100;
  opacity: 0.55;
  animation-name: gear-fade-in;
  animation-duration: 0.75s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
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
p {
  font-size: 3rem;
  margin: 1rem 0;
}
a {
  color: inherit;
  text-decoration: underline;
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
.container {
  position: relative;
}
</style>
