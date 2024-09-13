<template>
  <div class="outer-container">
    <h1 v-intersection-observer="[onIntersectionObserver]">
      Welcome to HackNJIT!
    </h1>

    <main>
      <img
        src="../assets/Slides/hacknjit_1.jpg"
        class="inner_img"
        v-intersection-observer="[onIntersectionObserver]"
      />
      <p class="inner_text" v-intersection-observer="[onIntersectionObserver]">
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing. Students break into small teams of up to 4,
        work together over 24 hours to create a tech project, and at the end
        present it to our judges!
      </p>
    </main>
    <img
      id="gear1"
      class="floating-gear"
      src="../assets/HackNJIT2024/gears/gear1.svg"
    />
  </div>
</template>

<script setup>
import { vIntersectionObserver } from "@vueuse/components";

function onIntersectionObserver([{ isIntersecting, target }]) {
  if (isIntersecting) {
    target.classList.add("fade-in");
  }
}
</script>

<style scoped>
.floating-gear {
  position: absolute;
  top: 0%;
  left: -20%;
  z-index: -100;
  width: 100vw;

  border: none;
  opacity: 0.55;
  animation-name: gear-fade-in;
  animation-duration: 0.75s;
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
.outer-container {
  position: relative;
  width: 60%;
  margin: 0rem auto;
  margin-bottom: 5rem;
  /* color: var(--color3); */
}
h1 {
  font-size: 4rem;
  margin: 1.75rem 0;
  font-weight: bold;
}
main {
  display: flex;
  gap: 2rem;
  justify-content: space-between;
  align-content: center;
}
p {
  flex: 1;
  font-size: 2rem;
  height: fit-content;
  align-self: center;
  font-weight: bold;
  text-align: left;
}
img {
  max-width: 50%;
  /* border: 4px var(--color3) solid; */
  border: 4px white solid;
  border-radius: 30px;
  flex: 1;
}
.fade-in {
  animation-name: fade-in;
  animation-duration: 0.75s;
  /* animation-delay: 1.5s; */
  animation-iteration-count: 1;
  animation-timing-function: linear;
  animation-fill-mode: forwards;
}
.fade {
  animation: fade-in 0.5s linear 1 normal forwards;
}

@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

@keyframes gear-fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 0.55;
  }
}

@media (max-width: 1750px) {
  .outer-container {
    width: 70%;
  }
}
@media (max-width: 1500px) {
  .outer-container {
    width: 77.5%;
  }
}
@media (max-width: 1200px) {
  main {
    display: flex;
    flex-wrap: wrap;
  }
  img {
    min-width: 100%;
  }
  .inner_text {
    min-width: 60%;
  }
}
@media (max-width: 750px) {
  p {
    font-size: 1.75rem;
  }
}
@media (max-width: 500px) {
  h1 {
    font-size: 3rem;
  }
}
</style>