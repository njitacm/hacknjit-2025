<template>
  <div class="NotifSignup outer-container">
    <main>
      <h1>HackNJIT</h1>
      <p class="inner_text">
        HackNJIT is a 24-hour hackathon at the New Jersey Institute of
        Technology, run by its ACM student chapter in conjunction with the Ying
        Wu College of Computing.
      </p>
      <p>
        Stay tuned, we'll return in November 2025!
      </p>
      <a :href="vw <= modalNoShowThreshold ? 'https://docs.google.com/forms/d/e/1FAIpQLSfiGXXticr2s7PcrMUN69K0U8LWq5sM4bnRJf-H4pfGU6MUNg/viewform?usp=dialog' : null"
        target="_blank" class="email-signup">
        <button @click="openModal">
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
            frameborder="0" marginheight="0" marginwidth="0">Loading…</iframe>
        </template>
      </Modal>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeMount } from "vue";
import Modal from './Modal.vue';

const modalNoShowThreshold = 400;   // viewport width in px;
const getVw = () => Math.max(document.documentElement.clientWidth || 0, window.innerWidth || 0);
const vw = ref(getVw());

const onWindowResize = () => {
  vw.value = getVw();
}

onMounted(() => {
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
  padding: 12px 24px;
  border: none;
  font-weight: bold;
  font-size: 2rem;
  cursor: pointer;
}

iframe {
  max-width: 100%;
  max-height: 100%;
  width: 750px;
  height: 550px;
  overflow: hidden;
}

.outer-container {
  display: grid;
  grid-template-rows: 1fr auto;
  gap: 32px;
  align-content: center;
  width: 95%;
  margin: 0rem auto;
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

main>* {
  height: min-content;
}

p {
  font-size: 2rem;
  height: -moz-fit-content;
  height: fit-content;
  align-self: center;
  text-align: center;
}

img {
  max-width: 50%;
  border-radius: 30px;
  flex: 1;
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

  .email-signup button {
    font-size: 1.5rem;
  }
}

@media(max-width: 650px) {
  iframe {
    height: 650px;
  }
}

@media (max-width: 450px) {
  h1 {
    font-size: 2.25rem;
  }

  .email-signup button {
    font-size: 1.25rem;
  }
}

/* .v-enter-active,
.v-leave-active {
  transition: all 1s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
} */
</style>