<template>
  <div class="outer-container">
    
    <main>
      <h1 v-intersection-observer="[onIntersectionObserver]">HackNJIT</h1>
      <!-- <img
        src="../assets/HackNJIT2024/clock_man.svg"
        class="inner_img"
        v-intersection-observer="[onIntersectionObserver]"
      /> -->
      <p class="inner_text" v-intersection-observer="[onIntersectionObserver]">
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing.
      </p>
      <p>
        Stay tuned, we'll return in November 2025!
      </p>
      <!-- <p class="inner_text" v-intersection-observer="[onIntersectionObserver]">
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing. Students break into small teams of up to 4,
        work together over 24 hours to create a tech project, and at the end
        present it to our judges!
      </p> -->
    </main>
    <div id="view-hint"><span>View last year's photos</span><span>&#x1F863;</span></div>
    <!-- <img
      id="gear1"
      class="floating-gear"
      src="../assets/HackNJIT2024/gears/gear1.svg"
    /> -->
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
  height: 100lvh;
  position: relative;
  width: 95%;
  margin: 0rem auto;
  margin-bottom: 2rem;
  /* color: var(--color3); */
}


h1 {
  font-size: 5rem;
  margin: 5rem 0 1.75rem 0;
  font-weight: bold;
}

main {
  display: grid;
  gap: 30px;
  width: 90%;
  max-width: 1500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -62.5%);
}

p {
  font-size: 2.5rem;
  height: -moz-fit-content;
  height: fit-content;
  align-self: center;
  text-align: center;
}

#view-hint {
  display: grid;
  gap: 10px;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
}

img {
  max-width: 50%;
  /* border: 4px var(--color3) solid; */
  /* border: 4px white solid; */
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
    /* width: 70%; */
  }
}

@media (max-width: 1550px) {
  .outer-container {
    /* width: 77.5%; */
  }

  /* p {
    font-size: 1.5rem;
  } */
}

@media (max-width: 1350px) {
  .outer-container {
    /* width: 77.5%; */
  }

  /* p {
    font-size: 1rem;
  } */
}

@media (max-width: 1200px) {
  main {
    display: flex;
    flex-wrap: wrap;
  }

  img {
    min-width: 55%;
    margin: 0 auto;
  }

  p {
    font-size: 1.5rem;
  }
}

@media (max-width: 1100px) {
  /* p {
    font-size: 1.5rem;
  } */
}

@media (max-width: 900px) {
  img {
    min-width: 50%;
    margin: 0 auto;
  }

  /* p {
    font-size: 1.25rem;
  } */
}

@media (max-width: 750px) {
  h1 {
    font-size: 3rem;
  }

  main {
    margin-top: 3rem;
  }

  /* p {
    font-size: 1.2rem;
  } */
}

@media (max-width: 700px) {
  main {
    flex-direction: column;
  }

  img {
    max-width: 75%;
  }

  p {
    /* font-size: 1.25rem; */
    text-align: center;
  }

  .outer-container {
    width: 90%;
  }
}

@media (max-width: 450px) {

  /* p {
    font-size: 1rem;
  } */
  h1 {
    font-size: 2.25rem;
  }
}
</style>