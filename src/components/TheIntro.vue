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
      <a :href="vw <= modalNoShowThreshold ? 'https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?usp=dialog' : null" target="_blank" class="email-signup">
        <button  @click="openModal">
          Sign up for email updates
        </button>
      </a>
      <Modal :is-open="isModalOpen && vw > modalNoShowThreshold" @modalClose="closeModal" name="email-sign-up">
        <template #title>
          <span>Receive Email Updates</span>
        </template>
        <template #body>
          <iframe
            src="https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?embedded=true"
            width="400" height="450" frameborder="0" marginheight="0" marginwidth="0">Loading…</iframe>
          <!-- <div>Hello</div> -->
        </template>
      </Modal>
    </main>
    <div id="view-hint" ref="viewHint">
      <span>View last year's photos</span>
      <br />
      <span>&#129123;</span>
    </div>
    <!-- <img
      id="gear1"
      class="floating-gear"
      src="../assets/HackNJIT2024/gears/gear1.svg"
    /> -->
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeMount } from "vue";
import Modal from './Modal.vue';

const modalNoShowThreshold = 400;   // viewport width in px;
const vw = ref(Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0));

const onWindowResize = () => {
  vw.value = Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0);
}

onMounted(() =>  {
  window.addEventListener("resize", onWindowResize);
});

onBeforeMount(() => {
  window.removeEventListener("resize", onWindowResize);
});


const isModalOpen = ref(false);

const openModal = () => {
  isModalOpen.value = true;
}

const closeModal = () => {
  isModalOpen.value = false;
}

</script>

<style scoped>
.email-signup {
  border-radius: 100px;
}

.email-signup button {
  border-radius: 100px;
  background-color: white;
  color: #170800;
  padding: 12px 16px;
  border: none;
  font-weight: bold;
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
  padding-bottom: 32px;
  display: grid;
  grid-template-rows: 1fr auto;
  gap: 32px;
  align-content: center;
  height: 100svh;
  width: 95%;
  margin: 0rem auto;
  margin-bottom: 2rem;
}


h1 {
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

main>* {
  height: min-content;
}

p {
  height: -moz-fit-content;
  height: fit-content;
  align-self: center;
  text-align: center;
}

#view-hint {
  align-content: center;
}

#view-hint * {
  margin: 0 auto;
}

img {
  max-width: 50%;
  border-radius: 30px;
  flex: 1;
}

@keyframes gear-fade-in {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 0.55;
  }
}

</style>