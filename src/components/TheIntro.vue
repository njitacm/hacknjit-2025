<template>
  <div class="outer-container">

    <main>
      <h1>HackNJIT</h1>
      <!-- <img
        src="../assets/HackNJIT2024/clock_man.svg"
        class="inner_img"
        v-intersection-observer="[onIntersectionObserver]"
      /> -->
      <p class="inner_text">
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
      <button class="email-signup" @click="openModal">
        Sign up for email updates
      </button>
      <Modal :is-open="isModalOpen" @modalClose="closeModal" name="email-sign-up">
        <template #title>
          <span>Receive Email Updates</span>
        </template>
        <template #body>
          <iframe
            src="https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?embedded=true"
            width="640" height="450" frameborder="0" marginheight="0" marginwidth="0">Loading…</iframe>
          <!-- <div>Hello</div> -->
        </template>
      </Modal>
    </main>
    <div id="view-hint" ref="viewHint">
      <span>View last year's photos</span>
      <br />
      <span>&#x1F863;</span>
    </div>
    <!-- <img
      id="gear1"
      class="floating-gear"
      src="../assets/HackNJIT2024/gears/gear1.svg"
    /> -->
  </div>
</template>

<script setup>
import { ref } from "vue";
import Modal from './Modal.vue';

const isModalOpen = ref(false);

const openModal = () => {
  isModalOpen.value = true;
  console.log("modal opened");
}

const closeModal = () => {
  isModalOpen.value = false;
  console.log("modal closed");
}

</script>

<style scoped>
button.email-signup {
  border-radius: 100px;
  background-color: white;
  color: #170800;
  padding: 12px 16px;
  border: none;
  font-weight: bold;
  font-size: 2rem;
  cursor: pointer;
}

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
  display: grid;
  grid-template-rows: 1fr auto;
  gap: 32px;
  align-content: center;
  height: 100lvh;
  width: 95%;
  margin: 0rem auto;
  margin-bottom: 2rem;
}


h1 {
  font-size: 5rem;
  margin: 5rem 0 1.75rem 0;
  font-weight: bold;
}

main {
  margin: 0 auto;
  display: grid;
  gap: 30px;
  width: 90%;
  max-width: 1500px;
  justify-items: center;
  align-content: center; 
  align-items: center;
}

main > * {
  height: min-content;
}

p {
  font-size: 2rem;
  height: -moz-fit-content;
  height: fit-content;
  align-self: center;
  text-align: center;
}

#view-hint {
  align-content: center;
  /* display: grid; */
  /* gap: 10px; */
  /* position: absolute; */
  /* bottom: 0; */
  /* left: 50%; */
  /* transform: translateX(-50%); */
  /* opacity: 0; */
}

#view-hint * {
  margin: 0 auto;
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

@media (max-width: 1200px) {
  p {
    font-size: 1.5rem;
  }
}

@media (max-width: 750px) {
  h1 {
    font-size: 3rem;
  }
}

@media (max-width: 450px) {
  h1 {
    font-size: 2.25rem;
  }
}
</style>